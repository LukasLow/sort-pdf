# Changelog

## [v0.4.0] - 2026-05-30

### Sicherheit
- Filename-Sanitization: `/`, `\`, `:`, `\0`, `..` werden aus User-Input entfernt
- Symlink-Traversal im PDF-Server durch `filepath.EvalSymlinks` verhindert
- Mutex für Config-Lese-/Schreibzugriffe (kein Datenverlust bei Race-Conditions)
- RWMutex für thread-sicheren Zugriff auf `currentWorkDir`

### Fehlerbehandlung
- `SelectWorkingDirectory` hat jetzt try/catch (keine unhandled Promise Rejection)
- `moveFile` hat Fallback auf Copy+Delete bei cross-device Rename
- `conflictAction` wird in `handleMoveToSystemTrash` validiert
- Null-Checks auf `conflict` im MoveConflictDialog

### UX
- Wizard erlaubt wieder Weiter zu Actions ohne Korrespondent (für Todo/Trash)
- Seitencount zeigt "1 Seite" / "N Seiten" statt "N s"
- Loading-Status im PdfViewer wartet auf tatsächliches Embed-Ladeereignis

### Bereinigung
- `WritePDFTags` (dead code) entfernt
- `.textLayer` CSS (von pdf.js übrig, wird von `<embed>` nicht genutzt) entfernt

## [v0.3.0] - 2026-05-30

- Bei Dateikonflikt: Bestehende Datei wird in den macOS-System-Papierkorb gelegt,
  statt überschrieben – die aktuelle Datei wird dann normal verschoben

## [v0.2.2] - 2026-05-30

- Wizard wird nach erfolgreichem Verschieben zurück auf View 1 gesetzt (neue PDF)
- Fehler beim Verschieben zeigen jetzt einen Alert mit Fehlermeldung

## [v0.2.1] - 2026-05-30

- Blaue Custom-Scrollbar entfernt (native macOS-Scrollbar wird verwendet)

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
