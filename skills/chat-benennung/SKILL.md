---
name: chat-benennung
description: Benenne den aktuellen Codex-Chat anhand einer vorhandenen Jira- oder Zammad-Ticketnummer und der konkreten Tätigkeit.
---

# Codex-Chat benennen

Dieser Skill ist Codex-spezifisch. Verwende eine verfügbare Funktion zum Umbenennen des aktuellen Chats. Fehlt diese Funktion oder unterstützt die Umgebung keine Chat-Benennung, benenne die Einschränkung und behaupte keine erfolgreiche Umbenennung. Ändere stattdessen keine Dateien oder andere Chats.

Benenne den aktuellen Codex-Chat, sobald sein Thema erkennbar ist. Verwende einen kurzen deutschen Titel mit einem Nomen für den Gegenstand und einem Verb im Infinitiv für die Tätigkeit.

- Ist eine zugehörige Jira-Ticketnummer bekannt, verwende `[<JiraTicketNr>] <Nomen> <Verb>`.
- Andernfalls, falls eine zugehörige Zammad-Ticketnummer bekannt ist, verwende `[<ZammadTicketNr>] <Nomen> <Verb>`.
- Ist keine Ticketnummer bekannt, verwende `<Nomen> <Verb>`.

Bevorzuge bei einem Chat mit beiden Nummern die Jira-Ticketnummer. Erfinde keine Nummer. Wenn später eine passende Ticketnummer oder ein genaueres Thema bekannt wird, aktualisiere den Titel. Benenne nur den aktuellen Chat; ändere bestehende andere Chats nur auf ausdrücklichen Wunsch.

Beispiele: `[ABC-1234] Mitarbeiterjournal öffnen`, `[661234] Mitarbeiterjournal öffnen`.
