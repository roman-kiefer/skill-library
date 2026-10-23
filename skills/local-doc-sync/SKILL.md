---
name: local-doc-sync
description: Aktualisiere den lokalen Dokumentationsindex auf ausdrueckliche Anfrage und melde Dokumentzahl, Fehler und fehlende Quellen.
license: See LICENSE
---

# Dokumentationsindex aktualisieren

Nutze diesen Skill, wenn der Nutzer den zentralen Dokumentationsindex aktualisieren oder neu einlesen lassen moechte. Normale Dokumentationsfragen gehoeren zum Skill `local-doc-search` und loesen hier keine Synchronisierung aus.

Verwende einen im Auftrag ausdrücklich genannten Projektpfad; andernfalls den Wert der Umgebungsvariable `LOCAL_DOCS_ROOT`. Fehlt beides, frage einmalig nach dem Pfad zum `local-docs`-Projekt. Leite den Projektpfad nicht aus dem Speicherort dieses Skills oder einer Skill-Verknüpfung ab.

Prüfe, dass der Pfad zum richtigen `local-docs`-Projekt mit passender `go.mod` gehört. Lies dort die geltende `AGENTS.md` und prüfe vor der Nutzung, dass `config.json` vorhanden ist. Fehlen Projekt oder Einrichtung, benenne die fehlende Voraussetzung; erfinde keine Pfade und lege keine Konfiguration ungefragt an. Führe die Go-Befehle aus diesem Projektverzeichnis aus.

Führe im Projektverzeichnis `go run . sync` aus. Verändere weder die Quellordner noch deren Originaldateien.

Pruefe danach die Ausgabe und `data/manifest.json`: Anzahl der Dokumente, Eintraege mit `status: error` und Warnungen zu fehlenden Quellordnern. Melde diese Zahlen sowie den Zeitpunkt der Aktualisierung. Bei Fehlern nenne die betroffenen Originalpfade und die Ursache. Ein nicht erreichbarer Quellordner bleibt im Manifest erhalten und darf nicht als leer interpretiert werden.

Die Synchronisierung prueft SHA-256 und extrahiert nur geaenderte Dokumente neu. Sie erstellt keinen Zeitplan.

## Herkunft und Lizenz

Übernommen und für die unabhängige Skill-Bibliothek angepasst aus [roman-kiefer/local-docs](https://github.com/roman-kiefer/local-docs). Die dort vorhandene GPLv3-Lizenzdatei wird unverändert als [LICENSE](LICENSE) mitgeführt.
