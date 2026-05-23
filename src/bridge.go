package src

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

	err = InitializeWorkDir(dir)
	if err != nil {
		return "", fmt.Errorf("fehler beim Erstellen der Ordner: %w", err)
	}

	err = SaveWorkDir(dir)
	if err != nil {
		return "", fmt.Errorf("fehler beim Speichern der Einstellungen: %w", err)
	}

	b.currentWorkDir = dir
	return dir, nil
}

// EnsureWorkDirStructure prüft, ob alle benötigten Ordner existieren,
// stellt fehlende wieder her und meldet, welche fehlten
func (b *WorkspaceBridge) EnsureWorkDirStructure() (string, error) {
	if b.currentWorkDir == "" {
		return "", fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}

	var missing []string
	for _, folder := range GetRequiredFolders() {
		path := filepath.Join(b.currentWorkDir, folder)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			missing = append(missing, folder)
		}
	}

	if len(missing) == 0 {
		return "", nil
	}

	err := InitializeWorkDir(b.currentWorkDir)
	if err != nil {
		return "", fmt.Errorf("fehler beim Wiederherstellen der Ordner: %w", err)
	}

	if len(missing) == 1 {
		return "Ordner " + missing[0] + " wurde wiederhergestellt.", nil
	}

	msg := "Folgende Ordner wurden wiederhergestellt: "
	for i, f := range missing {
		if i > 0 {
			msg += ", "
		}
		msg += f
	}
	return msg + ".", nil
}

// GetNextPDF liefert den Dateinamen der ältesten PDF im Eingang
func (b *WorkspaceBridge) GetNextPDF() (string, error) {
	if b.currentWorkDir == "" {
		return "", fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	return GetOldestInboxPDF(b.currentWorkDir)
}

// GetPdfInfo returns metadata for a PDF
func (b *WorkspaceBridge) GetPdfInfo(fileName string) (*PdfInfo, error) {
	if b.currentWorkDir == "" {
		return nil, fmt.Errorf("current work dir not set")
	}
	return getPdfInfo(b.currentWorkDir, fileName)
}

// --- Korrespondenten ---

func (b *WorkspaceBridge) GetCorrespondents() ([]string, error) {
	config, err := LoadFullConfig()
	if err != nil {
		return nil, err
	}
	return config.Correspondents, nil
}

func (b *WorkspaceBridge) AddCorrespondent(name string) error {
	config, err := LoadFullConfig()
	if err != nil {
		return err
	}
	for _, c := range config.Correspondents {
		if c == name {
			return nil
		}
	}
	config.Correspondents = append(config.Correspondents, name)
	return SaveFullConfig(config)
}

func (b *WorkspaceBridge) RemoveCorrespondent(name string) error {
	config, err := LoadFullConfig()
	if err != nil {
		return err
	}
	var updated []string
	for _, c := range config.Correspondents {
		if c != name {
			updated = append(updated, c)
		}
	}
	config.Correspondents = updated
	delete(config.CorrespondentFolders, name)
	return SaveFullConfig(config)
}

func (b *WorkspaceBridge) ReorderCorrespondents(reordered []string) error {
	config, err := LoadFullConfig()
	if err != nil {
		return err
	}
	config.Correspondents = reordered
	return SaveFullConfig(config)
}

// --- Korrespondenten-Ordner-Zuordnung ---

func (b *WorkspaceBridge) GetCorrespondentFolders() (map[string]string, error) {
	config, err := LoadFullConfig()
	if err != nil {
		return nil, err
	}
	return config.CorrespondentFolders, nil
}

func (b *WorkspaceBridge) SetCorrespondentFolder(correspondent string, folder string) error {
	config, err := LoadFullConfig()
	if err != nil {
		return err
	}
	if config.CorrespondentFolders == nil {
		config.CorrespondentFolders = map[string]string{}
	}
	if folder == "" {
		delete(config.CorrespondentFolders, correspondent)
	} else {
		config.CorrespondentFolders[correspondent] = folder
	}
	return SaveFullConfig(config)
}

// --- Ordner-Verwaltung (+Archiv) ---

func (b *WorkspaceBridge) GetFolderList() ([]string, error) {
	if b.currentWorkDir == "" {
		return nil, fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	archivDir := filepath.Join(b.currentWorkDir, "+Archiv")
	entries, err := ioutil.ReadDir(archivDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	var folders []string
	for _, e := range entries {
		if e.IsDir() {
			folders = append(folders, e.Name())
		}
	}
	return folders, nil
}

func (b *WorkspaceBridge) CreateFolder(name string) error {
	if b.currentWorkDir == "" {
		return fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	path := filepath.Join(b.currentWorkDir, "+Archiv", name)
	return os.MkdirAll(path, 0755)
}

// --- PDF-Analyse ---

// Analysis enthält die Ergebnisse der PDF-Analyse
type Analysis struct {
	DateYear      int    `json:"dateYear"`
	DateMonth     int    `json:"dateMonth"`
	DateDay       int    `json:"dateDay"`
	Correspondent string `json:"correspondent"`
}

// AnalyzePDF extrahiert Text, sucht nach Datum und Korrespondent
func (b *WorkspaceBridge) AnalyzePDF(fileName string) (*Analysis, error) {
	if b.currentWorkDir == "" {
		return nil, fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}

	path := filepath.Join(b.currentWorkDir, "900-Eingang", fileName)
	config, err := LoadFullConfig()
	if err != nil {
		return nil, err
	}

	return analyzePDF(path, config.Correspondents)
}

// WritePDFTags schreibt Tags in die PDF-Metadaten
func (b *WorkspaceBridge) WritePDFTags(fileName string, tags []string) error {
	if b.currentWorkDir == "" {
		return fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	path := filepath.Join(b.currentWorkDir, "900-Eingang", fileName)
	return writePDFTags(path, tags)
}

// --- Datei-Aktionen (mit Umbenennung) ---

func (b *WorkspaceBridge) moveFile(srcName string, targetDir string, targetName string) error {
	if b.currentWorkDir == "" {
		return fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	src := filepath.Join(b.currentWorkDir, "900-Eingang", srcName)
	dst := filepath.Join(targetDir, targetName)
	return os.Rename(src, dst)
}

func (b *WorkspaceBridge) MoveToTodo(fileName string, targetName string) error {
	target := filepath.Join(b.currentWorkDir, "800-TODO")
	return b.moveFile(fileName, target, targetName)
}

func (b *WorkspaceBridge) MoveToTrash(fileName string, targetName string) error {
	target := filepath.Join(b.currentWorkDir, "999-Trash")
	return b.moveFile(fileName, target, targetName)
}

func (b *WorkspaceBridge) MoveToArchiv(fileName string, subFolder string, targetName string) error {
	target := filepath.Join(b.currentWorkDir, "+Archiv", subFolder)
	return b.moveFile(fileName, target, targetName)
}
