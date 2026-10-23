// Read one Zammad ticket and its articles using environment credentials.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type config struct {
	base  string
	token string
}

type result struct {
	Ticket   map[string]json.RawMessage   `json:"ticket"`
	Articles []map[string]json.RawMessage `json:"articles"`
}

func configuration(lookup func(string) string) (config, error) {
	base := strings.TrimRight(strings.TrimSpace(lookup("ZAMMAD_URL")), "/")
	token := strings.TrimSpace(lookup("ZAMMAD_API_TOKEN"))
	if base == "" || token == "" {
		return config{}, errors.New("ZAMMAD_URL und ZAMMAD_API_TOKEN müssen gesetzt sein.")
	}
	u, err := url.Parse(base)
	if err != nil || u.Scheme != "https" || u.Hostname() == "" || u.User != nil || u.Path != "" || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" || strings.ContainsAny(base, " \t\r\n") {
		return config{}, errors.New("ZAMMAD_URL muss eine HTTPS-Startadresse ohne Pfad sein.")
	}
	for _, char := range token {
		if char < 33 || char > 126 {
			return config{}, errors.New("ZAMMAD_API_TOKEN enthält ungültige Zeichen.")
		}
	}
	return config{base: base, token: token}, nil
}

func noRedirects(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}

func get(client *http.Client, cfg config, path string, target any) error {
	req, err := http.NewRequest(http.MethodGet, cfg.base+"/api/v1/"+path, nil)
	if err != nil {
		return errors.New("Ungültige Zammad-Konfiguration.")
	}
	req.Header.Set("Authorization", "Token token="+cfg.token)
	req.Header.Set("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		// Raw transport errors may contain URLs or other confidential values.
		return errors.New("Zammad konnte nicht sicher erreicht werden (Netzwerk, TLS oder Zeitlimit).")
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		message := fmt.Sprintf("Zammad meldet HTTP %d.", response.StatusCode)
		switch {
		case response.StatusCode >= 300 && response.StatusCode < 400:
			message += " Weiterleitungen werden nicht verfolgt; Instanzadresse prüfen."
		case response.StatusCode == 401 || response.StatusCode == 403:
			message += " Token und Berechtigungen lokal prüfen."
		case response.StatusCode == 404:
			message += " Interne Ticket-ID und Instanz prüfen."
		}
		return errors.New(message)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.New("Die Zammad-Antwort konnte nicht vollständig gelesen werden.")
	}
	if err := json.Unmarshal(body, target); err != nil {
		return errors.New("Zammad hat ein ungültiges Datenformat geliefert.")
	}
	return nil
}

func readTicket(client *http.Client, cfg config, id string) (result, error) {
	var data result
	if err := get(client, cfg, "tickets/"+id, &data.Ticket); err != nil {
		return result{}, err
	}
	if data.Ticket == nil {
		return result{}, errors.New("Zammad hat kein Ticketobjekt geliefert.")
	}
	if err := get(client, cfg, "ticket_articles/by_ticket/"+id, &data.Articles); err != nil {
		return result{}, err
	}
	if data.Articles == nil {
		return result{}, errors.New("Zammad hat keine Artikelliste geliefert.")
	}
	return data, nil
}

func run(args []string, lookup func(string) string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("zammad-api", flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.Usage = func() {
		fmt.Fprintln(stderr, "Aufruf: zammad-api TICKET_ID\nLiest Ticket und Artikel. Erwartet die interne ID, nicht die Ticketnummer.\nUmgebung: ZAMMAD_URL und ZAMMAD_API_TOKEN.")
	}
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if flags.NArg() != 1 {
		flags.Usage()
		return 2
	}
	id := flags.Arg(0)
	valid := id != "" && strings.IndexFunc(id, func(r rune) bool { return r < '0' || r > '9' }) == -1
	number, err := strconv.ParseUint(id, 10, 64)
	if !valid || err != nil || number == 0 {
		fmt.Fprintln(stderr, "Die interne Ticket-ID muss eine positive Ganzzahl sein.")
		return 2
	}
	cfg, err := configuration(lookup)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	client := &http.Client{Timeout: 30 * time.Second, CheckRedirect: noRedirects}
	data, err := readTicket(client, cfg, strconv.FormatUint(number, 10))
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		fmt.Fprintln(stderr, "Die Ausgabe konnte nicht geschrieben werden.")
		return 1
	}
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Getenv, os.Stdout, os.Stderr))
}
