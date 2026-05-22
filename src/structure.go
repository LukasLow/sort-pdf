package src

import (
	"io/ioutil"
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

// CreateDefaultYaml erstellt eine leere korespondenten.yml, falls sie fehlt
func CreateDefaultYaml(baseDir string) error {
	yamlPath := filepath.Join(baseDir, "korespondeten.yml")

	// Prüfen, ob Datei schon existiert
	if _, err := os.Stat(yamlPath); !os.IsNotExist(err) {
		return nil
	}

	// Ein Paar Beispiel-Daten als valides YAML
	initialContent := []byte("Firmen:\n  - Beispiel GmbH\n")
	return ioutil.WriteFile(yamlPath, initialContent, 0644)
}

// InitializeWorkDir bündelt die Erstellung aller Strukturen
func InitializeWorkDir(baseDir string) error {
	err := CreateSubfolders(baseDir)
	if err != nil {
		return err
	}

	return CreateDefaultYaml(baseDir)
}
