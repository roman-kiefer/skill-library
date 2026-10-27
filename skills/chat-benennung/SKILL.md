---
name: chat-benennung
description: Benenne den aktuellen Codex-Chat anhand einer vorhandenen Jira-Ticketnummer, GitHub-Issuenummer oder Zammad-Ticketnummer und der konkreten Tätigkeit. Bewahre bei automatischer Benennung bestehende Titel mit führendem [.
---

# Codex-Chat benennen

Dieser Skill ist Codex-spezifisch. Verwende eine verfügbare Funktion zum Umbenennen des aktuellen Chats. Fehlt diese Funktion oder unterstützt die Umgebung keine Chat-Benennung, benenne die Einschränkung und behaupte keine erfolgreiche Umbenennung. Ändere stattdessen keine Dateien oder andere Chats.

Prüfe vor jeder automatischen Umbenennung den aktuellen Chattitel. Beginnt er mit `[`, lasse den gesamten Titel unverändert, unabhängig davon, ob später ein genaueres Thema, eine andere Ticketnummer oder ein PR bekannt wird. Dies gilt für jeden Titel mit führendem `[`. Kann der aktuelle Titel nicht ermittelt werden, führe keine automatische Umbenennung durch. Eine ausdrücklich vom Nutzer gewünschte Umbenennung bleibt möglich.

Beginnt der aktuelle Titel nicht mit `[`, benenne den aktuellen Codex-Chat automatisch, sobald sein Thema erkennbar ist. Verwende einen kurzen deutschen Titel mit einem Nomen für den Gegenstand und einem Verb im Infinitiv für die Tätigkeit.

- Ist eine zugehörige Jira-Ticketnummer bekannt, verwende `[<JiraTicketNr>] <Nomen> <Verb>`.
- Andernfalls, falls eine zugehörige GitHub-Issuenummer bekannt ist, verwende `[#<GitHubIssueNr>] <Nomen> <Verb>`.
- Andernfalls, falls eine zugehörige Zammad-Ticketnummer bekannt ist, verwende `[<ZammadTicketNr>] <Nomen> <Verb>`.
- Ist keine Ticketnummer bekannt, verwende `<Nomen> <Verb>`.

Verwende nur bekannte, zum Thema gehörende Nummern. Bei mehreren zugehörigen Nummern gilt die Priorität Jira → GitHub → Zammad. Erfinde keine Nummer. Wenn später eine passende Jira-Ticketnummer, GitHub-Issuenummer, Zammad-Ticketnummer oder ein genaueres Thema bekannt wird, aktualisiere den Titel nur, wenn der aktuelle Titel ermittelt wurde und nicht mit `[` beginnt, oder der Nutzer die Umbenennung ausdrücklich wünscht. Beachte dabei die genannte Priorität. Benenne nur den aktuellen Chat; ändere bestehende andere Chats nur auf ausdrücklichen Wunsch.

Beispiele: `[ABC-1234] Mitarbeiterjournal öffnen`, `[#123] Anmeldung reparieren`, `[661234] Mitarbeiterjournal öffnen`.
