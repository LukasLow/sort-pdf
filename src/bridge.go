package src

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type WorkspaceBridge struct {
	ctx            context.Context
	mu             sync.RWMutex
	currentWorkDir string
}

func NewWorkspaceBridge() *WorkspaceBridge {
	return &WorkspaceBridge{}
}

func (b *WorkspaceBridge) SetContext(ctx context.Context) {
	b.ctx = ctx
}

func (b *WorkspaceBridge) getWorkDir() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.currentWorkDir
}

func (b *WorkspaceBridge) setWorkDir(dir string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.currentWorkDir = dir
}

// CheckInitialWorkDir wird beim Start gerufen: Gibt es schon einen Ordner?
func (b *WorkspaceBridge) CheckInitialWorkDir() string {
	savedDir, err := LoadSavedWorkDir()
	if err != nil {
		return ""
	}
	b.setWorkDir(savedDir)
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

	b.setWorkDir(dir)
	return dir, nil
}

// EnsureWorkDirStructure prüft, ob alle benötigten Ordner existieren,
// stellt fehlende wieder her und meldet, welche fehlten
func (b *WorkspaceBridge) EnsureWorkDirStructure() (string, error) {
	if b.getWorkDir() == "" {
		return "", fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}

	var missing []string
	for _, folder := range GetRequiredFolders() {
		path := filepath.Join(b.getWorkDir(), folder)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			missing = append(missing, folder)
		}
	}

	if len(missing) == 0 {
		return "", nil
	}

	err := InitializeWorkDir(b.getWorkDir())
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
	if b.getWorkDir() == "" {
		return "", fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	return GetOldestInboxPDF(b.getWorkDir())
}

// GetPdfInfo returns metadata for a PDF
func (b *WorkspaceBridge) GetPdfInfo(fileName string) (*PdfInfo, error) {
	if b.getWorkDir() == "" {
		return nil, fmt.Errorf("current work dir not set")
	}
	return getPdfInfo(b.getWorkDir(), fileName)
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
	if b.getWorkDir() == "" {
		return nil, fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	archivDir := filepath.Join(b.getWorkDir(), "+Archiv")
	entries, err := os.ReadDir(archivDir)
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
	if b.getWorkDir() == "" {
		return fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	path := filepath.Join(b.getWorkDir(), "+Archiv", name)
	return os.MkdirAll(path, 0755)
}

// --- PDF-Analyse ---

// Analysis enthält die Ergebnisse der PDF-Analyse
type Analysis struct {
	Correspondent string `json:"correspondent"`
}

// AnalyzePDF extrahiert Text und sucht nach Korrespondent
func (b *WorkspaceBridge) AnalyzePDF(fileName string) (*Analysis, error) {
	if b.getWorkDir() == "" {
		return nil, fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}

	path := filepath.Join(b.getWorkDir(), "900-Eingang", fileName)
	config, err := LoadFullConfig()
	if err != nil {
		return nil, err
	}

	return analyzePDF(path, config.Correspondents)
}

// GetPDFPageCount gibt die Anzahl der Seiten einer PDF zurück
func (b *WorkspaceBridge) GetPDFPageCount(fileName string) (int, error) {
	if b.getWorkDir() == "" {
		return 0, fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	path := filepath.Join(b.getWorkDir(), "900-Eingang", fileName)
	return getPDFPageCount(path)
}

// --- Datei-Aktionen (mit Umbenennung) ---

func (b *WorkspaceBridge) moveFile(srcName string, targetDir string, targetName string) error {
	if b.getWorkDir() == "" {
		return fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	src := filepath.Join(b.getWorkDir(), "900-Eingang", srcName)
	dst := filepath.Join(targetDir, targetName)
	if err := os.Rename(src, dst); err != nil {
		// Fallback bei cross-device link: kopieren + löschen
		var linkErr *os.LinkError
		if errors.As(err, &linkErr) {
			if copyErr := b.copyFile(src, dst); copyErr != nil {
				return fmt.Errorf("rename fehlgeschlagen (cross-device), copy ebenfalls: %w", copyErr)
			}
			if removeErr := os.Remove(src); removeErr != nil {
				return fmt.Errorf("datei kopiert, aber Quelle konnte nicht gelöscht werden: %w", removeErr)
			}
		} else {
			return err
		}
	}
	inboxPath := filepath.Join(b.getWorkDir(), "900-Eingang")
	trashPath := filepath.Join(b.getWorkDir(), "999-Trash")
	b.cleanupParentDirs(inboxPath, trashPath, src)
	return nil
}

func (b *WorkspaceBridge) copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// cleanupParentDirs entfernt leere Eltern-Ordner einer verschobenen Datei
// und verschiebt sie rekursiv nach 999-Trash.
func (b *WorkspaceBridge) cleanupParentDirs(inboxPath, trashPath, movedFilePath string) {
	dir := filepath.Dir(movedFilePath)
	for {
		if dir == inboxPath {
			break
		}
		if !strings.HasPrefix(dir, inboxPath) {
			break
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			log.Printf("cleanupParentDirs: kann Ordner nicht lesen %s: %v", dir, err)
			return
		}
		if len(entries) > 0 {
			break
		}
		trashName := filepath.Base(dir) + "_" + time.Now().Format("20060102150405")
		trashDst := filepath.Join(trashPath, trashName)
		if err := os.Rename(dir, trashDst); err != nil {
			log.Printf("cleanupParentDirs: kann Ordner nicht verschieben %s -> %s: %v", dir, trashDst, err)
			return
		}
		dir = filepath.Dir(dir)
	}
}

// MoveConflict describes a naming conflict when a file already exists at the target.
type MoveConflict struct {
	HasConflict bool   `json:"hasConflict"`
	SourcePath  string `json:"sourcePath"`
	TargetPath  string `json:"targetPath"`
	TargetName  string `json:"targetName"`
}

func (b *WorkspaceBridge) checkConflict(fileName, targetPath, targetName string) (*MoveConflict, error) {
	_, err := os.Stat(targetPath)
	if os.IsNotExist(err) {
		return &MoveConflict{HasConflict: false}, nil
	}
	if err != nil {
		return nil, err
	}
	srcPath := filepath.Join(b.getWorkDir(), "900-Eingang", fileName)
	return &MoveConflict{
		HasConflict: true,
		SourcePath:  srcPath,
		TargetPath:  targetPath,
		TargetName:  targetName,
	}, nil
}

// CheckArchiveConflict prüft, ob im +Archiv bereits eine Datei mit dem Zielnamen existiert.
func (b *WorkspaceBridge) CheckArchiveConflict(fileName, subFolder, targetName string) (*MoveConflict, error) {
	if b.getWorkDir() == "" {
		return nil, fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	target := filepath.Join(b.getWorkDir(), "+Archiv", subFolder, targetName)
	return b.checkConflict(fileName, target, targetName)
}

// CheckTodoConflict prüft, ob im 800-TODO bereits eine Datei mit dem Zielnamen existiert.
func (b *WorkspaceBridge) CheckTodoConflict(fileName, targetName string) (*MoveConflict, error) {
	if b.getWorkDir() == "" {
		return nil, fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	target := filepath.Join(b.getWorkDir(), "800-TODO", targetName)
	return b.checkConflict(fileName, target, targetName)
}

// CheckTrashConflict prüft, ob im 999-Trash bereits eine Datei mit dem Zielnamen existiert.
func (b *WorkspaceBridge) CheckTrashConflict(fileName, targetName string) (*MoveConflict, error) {
	if b.getWorkDir() == "" {
		return nil, fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	target := filepath.Join(b.getWorkDir(), "999-Trash", targetName)
	return b.checkConflict(fileName, target, targetName)
}

// OpenFileInOSViewer öffnet eine Datei mit dem systemeigenen Standardprogramm.
func (b *WorkspaceBridge) OpenFileInOSViewer(path string) error {
	return exec.Command("open", path).Start()
}

// MoveFileToSystemTrash verschiebt eine Datei in den macOS-System-Papierkorb.
func (b *WorkspaceBridge) MoveFileToSystemTrash(path string) error {
	escaped := strings.ReplaceAll(path, `\`, `\\`)
	escaped = strings.ReplaceAll(escaped, `"`, `\"`)
	cmd := exec.Command("osascript", "-e",
		fmt.Sprintf(`tell application "Finder" to delete POSIX file "%s"`, escaped))
	return cmd.Run()
}

// MoveToArchiv verschiebt die Datei ins +Archiv (schlägt fehl, wenn das Ziel existiert).
func (b *WorkspaceBridge) MoveToArchiv(fileName string, subFolder string, targetName string) error {
	target := filepath.Join(b.getWorkDir(), "+Archiv", subFolder)
	return b.moveFile(fileName, target, targetName)
}

// MoveToTodo verschiebt die Datei ins 800-TODO.
func (b *WorkspaceBridge) MoveToTodo(fileName string, targetName string) error {
	target := filepath.Join(b.getWorkDir(), "800-TODO")
	return b.moveFile(fileName, target, targetName)
}

// MoveToTrash verschiebt die Datei ins 999-Trash.
func (b *WorkspaceBridge) MoveToTrash(fileName string, targetName string) error {
	target := filepath.Join(b.getWorkDir(), "999-Trash")
	return b.moveFile(fileName, target, targetName)
}
