package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadTicket(t *testing.T) {
	var paths []string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		if r.Method != http.MethodGet || r.Header.Get("Authorization") != "Token token=test-secret" || r.Header.Get("Accept") != "application/json" {
			t.Error("unexpected request method or headers")
		}
		if r.URL.Path == "/api/v1/tickets/123" {
			io.WriteString(w, `{"id":123,"title":"Prüfung","large":9007199254740993}`)
		} else if r.URL.Path == "/api/v1/ticket_articles/by_ticket/123" {
			io.WriteString(w, `[{"id":7,"internal":true}]`)
		} else {
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	client := server.Client()
	client.CheckRedirect = noRedirects
	data, err := readTicket(client, config{server.URL, "test-secret"}, "123")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(paths, ",") != "/api/v1/tickets/123,/api/v1/ticket_articles/by_ticket/123" {
		t.Fatalf("unexpected requests: %v", paths)
	}
	if string(data.Ticket["large"]) != "9007199254740993" || string(data.Articles[0]["internal"]) != "true" {
		t.Fatal("ticket precision or internal visibility lost")
	}
	encoded, err := json.Marshal(data)
	if err != nil || bytes.Contains(encoded, []byte("test-secret")) {
		t.Fatal("invalid output or credential disclosure")
	}
}

func TestErrorsAndRedirects(t *testing.T) {
	for _, scenario := range []struct {
		name, ticket, articles, message string
		status                          int
	}{
		{"unauthorized", "test-secret", "[]", "401", 401},
		{"redirect", "test-secret", "[]", "Weiterleitungen", 302},
		{"invalid-json", "test-secret", "[]", "Datenformat", 200},
		{"null-ticket", "null", "[]", "Ticketobjekt", 200},
		{"null-articles", "{}", "null", "Artikelliste", 200},
		{"wrong-articles", "{}", "{}", "Datenformat", 200},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			redirected := false
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/other" {
					redirected = true
				}
				if r.URL.Path == "/api/v1/tickets/123" {
					w.Header().Set("Location", "/other")
					w.WriteHeader(scenario.status)
					io.WriteString(w, scenario.ticket)
				} else {
					io.WriteString(w, scenario.articles)
				}
			}))
			defer server.Close()
			client := server.Client()
			client.CheckRedirect = noRedirects
			_, err := readTicket(client, config{server.URL, "test-secret"}, "123")
			if err == nil || !strings.Contains(err.Error(), scenario.message) || strings.Contains(err.Error(), "test-secret") || redirected {
				t.Fatalf("unexpected error or redirect: %v, redirected=%v", err, redirected)
			}
		})
	}
}

func TestConfiguration(t *testing.T) {
	for _, base := range []string{"", "http://support.example.org", "https://user:pass@support.example.org", "https://support.example.org/api/v1", "https://support.example.org?key=value", "https://support.example.org/#ticket"} {
		_, err := configuration(func(key string) string {
			if key == "ZAMMAD_URL" {
				return base
			}
			return "test-secret"
		})
		if err == nil || strings.Contains(err.Error(), "test-secret") {
			t.Fatalf("unsafe configuration accepted: %q", base)
		}
	}
	cfg, err := configuration(func(key string) string {
		if key == "ZAMMAD_URL" {
			return "https://support.example.org/"
		}
		return "test-secret"
	})
	if err != nil || cfg.base != "https://support.example.org" {
		t.Fatalf("valid configuration rejected: %v", err)
	}
}

func TestCLI(t *testing.T) {
	for _, scenario := range []struct {
		args []string
		code int
	}{
		{[]string{"--help"}, 0},
		{nil, 2},
		{[]string{"0"}, 2},
		{[]string{"+1"}, 2},
		{[]string{"12/3"}, 2},
		{[]string{"123", "456"}, 2},
		{[]string{"123"}, 1},
		{[]string{"--number", "663140"}, 1},
		{[]string{"--number", "663140", "3144"}, 2},
		{[]string{"--number", "0"}, 2},
		{[]string{"--number", "663140 OR *"}, 2},
	} {
		var out, err bytes.Buffer
		code := run(scenario.args, func(string) string { return "" }, &out, &err)
		if code != scenario.code || out.Len() != 0 {
			t.Fatalf("args %v: code=%d, output=%s", scenario.args, code, out.String())
		}
	}
}

func TestResolveNumber(t *testing.T) {
	for _, scenario := range []struct {
		name, body, want string
		failure          bool
	}{
		{"exact-only", `[{"id":77,"number":"6631400"},{"id":3144,"number":"663140"}]`, "3144", false},
		{"duplicate-result", `[{"id":3144,"number":"663140"},{"id":3144,"number":"663140"}]`, "3144", false},
		{"not-found", `[]`, "", true},
		{"similar-only", `[{"id":77,"number":"6631400"}]`, "", true},
		{"ambiguous", `[{"id":1,"number":"663140"},{"id":2,"number":"663140"}]`, "", true},
		{"invalid-id", `[{"id":0,"number":"663140"}]`, "", true},
		{"null", `null`, "", true},
		{"malformed", `{}`, "", true},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" || r.URL.Path != "/api/v1/tickets/search" || r.URL.Query().Get("query") != "number:663140" || r.Header.Get("Authorization") != "Token token=test-secret" {
					t.Error("unexpected search request")
				}
				io.WriteString(w, scenario.body)
			}))
			defer server.Close()
			id, err := resolveNumber(server.Client(), config{server.URL, "test-secret"}, "663140")
			if (err != nil) != scenario.failure || id != scenario.want {
				t.Fatalf("id=%q err=%v", id, err)
			}
		})
	}
}

func TestResolveNumberPagination(t *testing.T) {
	for _, ambiguous := range []bool{false, true} {
		pages := 0
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pages++
			if r.URL.Query().Get("page") == "1" {
				tickets := make([]map[string]any, 100)
				for i := range tickets {
					tickets[i] = map[string]any{"id": i + 1, "number": "other"}
				}
				if ambiguous {
					tickets[0] = map[string]any{"id": 1, "number": "663140"}
				}
				json.NewEncoder(w).Encode(tickets)
			} else {
				io.WriteString(w, `[{"id":3144,"number":"663140"}]`)
			}
		}))
		id, err := resolveNumber(server.Client(), config{server.URL, "test-secret"}, "663140")
		server.Close()
		if pages != 2 || (err != nil) != ambiguous || (!ambiguous && id != "3144") {
			t.Fatalf("pages=%d id=%q err=%v", pages, id, err)
		}
	}
}
