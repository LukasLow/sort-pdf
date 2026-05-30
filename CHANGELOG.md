# Changelog

## [v0.2.0] - 2026-05-30

### Sicherheit
- Path-Traversal-Schutz im PDF-Server: Nur Dateien innerhalb von `900-Eingang` werden ausgeliefert
- Config-Datenverlust bei korrupten JSON-Dateien verhindert (leere Config überschreibt nicht mehr bestehende Daten)

### Bugfixes
- Hartkodiertes Jahr `2026` in der Datumsauswahl durch dynamisches Jahr ersetzt
- Ungültige Daten (z. B. 31. Februar) werden automatisch korrigiert
- `Weiter`-Button im Formular-Wizard erfordert jetzt einen ausgewählten Korrespondenten
- Korrespondenten ohne Ordner-Zuordnung sind jetzt im Dropdown sichtbar (mit Hinweis)
- Stumm geschluckte Fehler in der PDF-Suche werden jetzt geloggt
- Vor dem Verschieben einer PDF wird jetzt geprüft, ob am Zielort bereits eine Datei mit demselben Namen existiert; bei Konflikt erscheint ein Dialog mit der Option, beide PDFs zu öffnen, zu vergleichen und ggf. zu überschreiben

### Sonstiges
- Unbenutzten Methoden-Receiver in `GetVersion` entfernt

## [v0.1.2] - 2026-05-24

- Neues Design-System aus `farben.md`: Dark Theme mit Blau als Primary-Farbe
- Logo aktualisiert (Blau, quadratisches SVG)
- `logo.svg` in assets integriert und im Sidebar-Header eingebaut
- Altes Logo (`logo-universal.png`) entfernt
- Ordnerstruktur bereinigt

## [v0.1.1] - 2026-05-24

- CI: Archive im Build-Job direkt gepackt (macOS .app im .zip, Windows .exe im .zip, Linux binary im .tar.gz)
- Windows-Build: NSIS-Installer entfernt, rohe .exe in Zip

## [v0.1.0] - 2026-05-24

- Erstveröffentlichung
- PDF-Eingang sortieren mit Korrespondenten-Erkennung
- macOS-, Windows- und Linux-Builds per GitHub CI
