# sort-pdf

PDF-Sortier-Tool – Eingangsscan-PDFs automatisch Korrespondenten zuordnen und archivieren.

## Development

```bash
wails dev
```

## Build

```bash
# Version aus src/version.go wird automatisch eingebettet
wails build
```

## Release

Einfach einen Tag pushen:

```bash
git tag v0.1.0
git push origin v0.1.0
```

GitHub CI baut dann macOS-, Windows- und Linux-Pakete.

## Version aktualisieren

Die Version wird **nur** in `src/version.go` geändert. Alle Stellen
(Info.plist, Windows-Version, Frontend-Settings, Build-Artifakte)
ziehen daraus.
