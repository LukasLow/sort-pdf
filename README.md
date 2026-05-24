# sort-pdf

PDF-Sortier-Tool – Eingangsscan-PDFs automatisch Korrespondenten zuordnen und archivieren.

## Development

```bash
wails dev
```

## Build

```bash
# Version aus src/version.go wird in wails.json synchronisiert
./scripts/sync-version.sh
wails build
```

> **Eine Quelle für die Version:** `src/version.go` – das Skript `scripts/sync-version.sh`
> schreibt sie vor dem Build in `wails.json`. Der CI-Workflow macht dasselbe.

## Release

Einfach einen Tag pushen:

```bash
git tag v0.1.2
git push origin v0.1.2
```

Aktuelle Version: **v0.1.2**

GitHub CI baut dann macOS-, Windows- und Linux-Pakete.

## Version aktualisieren

Die Version wird **nur** in `src/version.go` geändert. Alle Stellen
(Info.plist, Windows-Version, Frontend-Settings, Build-Artifakte)
ziehen daraus.
