package src

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type WorkspaceBridge struct {
	ctx            context.Context
	currentWorkDir string
}

func NewWorkspaceBridge() *WorkspaceBridge {
	return &WorkspaceBridge{}
}

func (b *WorkspaceBridge) SetContext(ctx context.Context) {
	b.ctx = ctx
}

// CheckInitialWorkDir wird beim Start gerufen: Gibt es schon einen Ordner?
func (b *WorkspaceBridge) CheckInitialWorkDir() string {
	savedDir, err := LoadSavedWorkDir()
	if err != nil {
		return ""
	}
	b.currentWorkDir = savedDir
	return savedDir
}

// SelectWorkingDirectory öffnet den macOS-Dialog und richtet alles ein
func (b *WorkspaceBridge) SelectWorkingDirectory() (string, error) {
	dir, err := runtime.OpenDirectoryDialog(b.ctx, runtime.OpenDialogOptions{
		Title: "Arbeitsordner für PDF-Sortierung wählen",
	})
	if err != nil {
		return "", err
	}
	if dir == "" {
		return "", fmt.Errorf("kein Ordner ausgewählt")
	}

	// Struktur anlegen
	err = InitializeWorkDir(dir)
	if err != nil {
		return "", fmt.Errorf("fehler beim Erstellen der Ordner: %w", err)
	}

	// Pfad dauerhaft in den macOS App-Einstellungen sichern
	err = SaveWorkDir(dir)
	if err != nil {
		return "", fmt.Errorf("fehler beim Speichern der Einstellungen: %w", err)
	}

	b.currentWorkDir = dir
	return dir, nil
}

// GetNextPDF liefert den Dateinamen der ältesten PDF im Eingang an das Frontend
func (b *WorkspaceBridge) GetNextPDF() (string, error) {
	if b.currentWorkDir == "" {
		return "", fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	return GetOldestInboxPDF(b.currentWorkDir)
}

// GetPdfInfo returns metadata for a PDF in the current work directory
func (b *WorkspaceBridge) GetPdfInfo(fileName string) (*PdfInfo, error) {
	if b.currentWorkDir == "" {
		return nil, fmt.Errorf("current work dir not set")
	}
	return getPdfInfo(b.currentWorkDir, fileName)
}
