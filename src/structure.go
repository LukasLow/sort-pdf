package src

import (
	"os"
	"path/filepath"
)

// GetRequiredFolders liefert die Liste der Ordnernamen, die wir brauchen
func GetRequiredFolders() []string {
	return []string{
		"+Archiv",
		"800-TODO",
		"900-Eingang",
		"999-Trash",
	}
}

// CreateSubfolders erstellt die Standard-Ordnerstruktur im Arbeitsverzeichnis
func CreateSubfolders(baseDir string) error {
	for _, folder := range GetRequiredFolders() {
		path := filepath.Join(baseDir, folder)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// InitializeWorkDir bündelt die Erstellung aller Strukturen
func InitializeWorkDir(baseDir string) error {
	return CreateSubfolders(baseDir)
}
