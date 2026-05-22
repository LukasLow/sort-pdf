package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// AppConfig hält die session-übergreifenden Einstellungen
type AppConfig struct {
	WorkDir string `json:"workDir"`
}

// GetConfigPath ermittelt den macOS-Pfad: ~/Library/Application Support/sort-pdf/config.json
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, "Library", "Application Support", "sort-pdf")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "config.json"), nil
}

// LoadSavedWorkDir liest den gespeicherten Ordnerpfad aus der JSON-Datei
func LoadSavedWorkDir() (string, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "", nil // Noch keine Config vorhanden
	}

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	var config AppConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return "", err
	}

	return config.WorkDir, nil
}

// SaveWorkDir speichert den gewählten Ordnerpfad dauerhaft ab
func SaveWorkDir(path string) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	config := AppConfig{WorkDir: path}
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, bytes, 0644)
}
