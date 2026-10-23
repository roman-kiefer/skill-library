# Skill Library

Persönliche Skill-Bibliothek für den Skill Manager. Jeder Skill liegt als echter
Ordner unter `skills/<name>/SKILL.md`; zusätzliche Dateien gehören zum jeweiligen
Skill. Installation und Updates erfolgen über das Git-Repository.

| Skill | Aufgabe | Voraussetzungen |
| --- | --- | --- |
| `chat-benennung` | Aktuellen Codex-Chat nach Tätigkeit und gegebenenfalls Ticketnummer benennen | Codex mit verfügbarer Funktion zur Chat-Benennung |
| `local-doc-search` | Lokalen Dokumentenindex durchsuchen und Originalpfade sowie Seiten/Zeilen zitieren | Eingerichtetes `local-docs`-Projekt |
| `local-doc-sync` | Dokumentenindex auf ausdrücklichen Wunsch aktualisieren | Eingerichtetes `local-docs`-Projekt |
| `zammad-antwort` | Kundenantwort als Entwurf zum manuellen Kopieren erstellen | Tickettext oder lesender Zammad-Zugriff |
| `zammad-jira-check` | Passende Jira-Tickets zu einem Problem finden | Jira-Suchzugriff, bei Ticketbezug auch Zammad-Kontext |
| `zammad-notiz` | Interne Zammad-Notiz entwerfen oder ausdrücklich beauftragt eintragen | Ticketkontext; zum Eintragen verlässlicher interner Schreibzugriff |

## Mit dem Skill Manager installieren

Den Skill Manager mit dieser Quelle starten:

```powershell
skill-manager --repo https://github.com/roman-kiefer/skill-library.git --source C:\pfad\skill-library
```

`C:\pfad\skill-library` durch den lokalen Checkout-Pfad ersetzen. Beim ersten Start
ist die URL auch interaktiv konfigurierbar. Bei einem vorhandenen Checkout muss
`origin` zur angegebenen URL passen. Für macOS/Linux entsprechend einen lokalen
Pfad verwenden.

In der TUI mit der Leertaste markieren, mit Tab das Ziel (Codex, Claude Code oder
beide) wählen und mit Enter installieren. Ziele sind `~/.agents/skills` und
`~/.claude/skills`. Der Manager verlinkt die vollständigen Skill-Ordner; bestehende
abweichende Dateien oder Ordner bleiben als Konflikt erhalten. Updates der Quelle
werden über die Links direkt sichtbar. Bereits anderweitig installierte Skills
bei Konflikten zuerst prüfen; diese Bibliothek ersetzt sie nicht automatisch.

Die textbasierten Dokumenten- und Zammad-Skills sind auch mit Claude Code nutzbar,
sofern die benötigten Werkzeuge und Zugriffe dort verfügbar sind. `agents/openai.yaml`
enthält optionale Codex-Metadaten. Chat-Benennung bleibt Codex-spezifisch; ohne
passende Funktion wird keine erfolgreiche Umbenennung behauptet.

## Lokale Dokumentation anbinden

Das Suchprogramm wird separat aus [roman-kiefer/local-docs](https://github.com/roman-kiefer/local-docs)
eingerichtet. Benötigt werden Go (Version gemäß dessen `go.mod`), ripgrep und für
PDF-Verarbeitung Poppler. LibreOffice ist dort nur für aktivierte alte DOC-Dateien
erforderlich. Die `config.json` mit Quellen und Werkzeugpfaden bleibt im Suchprojekt.

Den Projektpfad für die laufende Sitzung setzen:

```powershell
$env:LOCAL_DOCS_ROOT = 'C:\pfad\local-docs'
```

```sh
export LOCAL_DOCS_ROOT=/pfad/local-docs
```

Codex bzw. Claude Code muss den Wert im eigenen Prozess sehen; nach einer dauerhaften
Änderung der Benutzerumgebung das Programm gegebenenfalls neu starten. Alternativ
den Projektpfad ausdrücklich im Auftrag nennen. Ohne Pfad fragt der Skill einmalig
nach. Der Speicherort eines installierten oder verlinkten Skills bestimmt das
Suchprojekt nicht.

Vor der Nutzung werden die Projektvorgaben und die vorhandene Konfiguration geprüft.
Originaldokumente bleiben unverändert. Suchtreffer zitieren Originalpfad und
PDF-Seite beziehungsweise Textzeile; die Synchronisierung erstellt keinen Zeitplan.
Suchprogramm, lokale Konfiguration, Dokumente und Indexdaten sind nicht Teil dieser
Bibliothek.

## Zammad und Jira

Es ist keine Instanzadresse fest hinterlegt. Ticketlinks, Auftrag und verfügbare
Anbindungen bestimmen den Kontext. Bei fehlender Anmeldung oder mehrdeutiger
Instanz wird gezielt nachgefragt. Für Entwürfe kann auch bereitgestellter Tickettext
genügen; eine tatsächliche Jira-Suche benötigt eine verfügbare Jira-Anbindung.

`zammad-antwort` liefert ausschließlich einen Entwurf. `zammad-notiz` liefert
standardmäßig ebenfalls einen Entwurf; nur ein ausdrücklicher Auftrag erlaubt das
Eintragen einer zuverlässig als intern geprüften Notiz. Ohne geeigneten Zugriff
bleibt es beim Entwurf. Kopierbare Texte funktionieren auch ohne ChatGPT-Schreibblöcke.

`zammad-jira-check` sucht ohne Ticketänderungen. Bei einem passenden Vorgang fragt
der Skill vor einer internen Notiz nach und verwendet anschließend `zammad-notiz`.
Für diesen vollständigen Ablauf beide Skills installieren; fehlt der Notiz-Skill,
wird höchstens ein Entwurf ausgegeben. Jira-interne Inhalte benötigen weiterhin eine
eigene konkrete Freigabe zur Übertragung; eine Notizfreigabe allein genügt nicht.

Die Bibliothek richtet keine Konten, Anbindungen oder Zugangsdaten ein. Es werden
keine echten Tickets oder Dokumente als Beispiele mitgeliefert.

## Herkunft und Lizenzen

Die Dokumenten-Skills stammen aus `roman-kiefer/local-docs` und wurden für einen
unabhängigen Installationsort angepasst. Die dort vorhandene GPLv3-Lizenzdatei ist
unverändert bei [local-doc-search](skills/local-doc-search/LICENSE) und
[local-doc-sync](skills/local-doc-sync/LICENSE) enthalten. Weitere lokale Skills
wurden als Ausgangspunkt übernommen und die Zammad-Anweisungen neutralisiert.
Die gesamte Bibliothek steht unter der GNU General Public License Version 3
(GPLv3). Den vollständigen Lizenztext enthält [LICENSE](LICENSE).
