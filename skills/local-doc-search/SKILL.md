---
name: local-doc-search
description: Durchsuche lokal indexierte Dokumente nach Fachfragen; liefere Originalpfade und PDF-Seiten.
license: See LICENSE
---

# Local document search

Verwende einen im Auftrag ausdrücklich genannten Projektpfad; andernfalls den Wert der Umgebungsvariable `LOCAL_DOCS_ROOT`. Fehlt beides, frage einmalig nach dem Pfad zum `local-docs`-Projekt. Leite den Projektpfad nicht aus dem Speicherort dieses Skills oder einer Skill-Verknüpfung ab.

Prüfe, dass der Pfad zum richtigen `local-docs`-Projekt mit passender `go.mod` gehört. Lies dort die geltende `AGENTS.md` und prüfe vor der Nutzung, dass `config.json` vorhanden ist. Fehlen Projekt oder Einrichtung, benenne die fehlende Voraussetzung; erfinde keine Pfade und lege keine Konfiguration ungefragt an. Führe die Go-Befehle aus diesem Projektverzeichnis aus.

Use `go run . search --query "phrase"` for focused research. Run `go run . sync` first only if the index is missing or the user requests current source content. Search output includes a document ID, version, original path, and PDF page number. Cite the original path and page in the answer. For visual inspection, run `go run . page --id ID --page N` and inspect the returned PNG. DOCX and ODT results have line references only.

The default search prefers `-nokey.pdf` within the same document family and version. Run `search --all` when fallback variants or differences between variants matter. Do not infer that similarly named documents from different versions contain the same content.

Original documents are read-only. Machine-specific source and tool paths live in the untracked `config.json`; generated index files live in `data/`.

## Herkunft und Lizenz

Übernommen und für die unabhängige Skill-Bibliothek angepasst aus [roman-kiefer/local-docs](https://github.com/roman-kiefer/local-docs). Die dort vorhandene GPLv3-Lizenzdatei wird unverändert als [LICENSE](LICENSE) mitgeführt.
