package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// AppConfig hält die session-übergreifenden Einstellungen
type AppConfig struct {
	WorkDir              string            `json:"workDir"`
	Correspondents       []string          `json:"correspondents"`
	CorrespondentFolders map[string]string `json:"correspondentFolders"`
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

// LoadFullConfig liest die gesamte Konfiguration aus der JSON-Datei
func LoadFullConfig() (AppConfig, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return AppConfig{}, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return AppConfig{Correspondents: []string{}, CorrespondentFolders: map[string]string{}}, nil
	}

	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return AppConfig{}, err
	}

	var config AppConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return AppConfig{}, err
	}

	if config.Correspondents == nil {
		config.Correspondents = []string{}
	}
	if config.CorrespondentFolders == nil {
		config.CorrespondentFolders = map[string]string{}
	}

	return config, nil
}

// SaveFullConfig speichert die gesamte Konfiguration in die JSON-Datei
func SaveFullConfig(config AppConfig) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, bytes, 0644)
}

// LoadSavedWorkDir liest den gespeicherten Ordnerpfad aus der JSON-Datei
func LoadSavedWorkDir() (string, error) {
	config, err := LoadFullConfig()
	if err != nil {
		return "", err
	}
	return config.WorkDir, nil
}

// SaveWorkDir speichert den gewählten Ordnerpfad dauerhaft ab
func SaveWorkDir(path string) error {
	config, err := LoadFullConfig()
	if err != nil {
		config = AppConfig{}
	}

	config.WorkDir = path

	if config.Correspondents == nil {
		config.Correspondents = []string{}
	}
	if config.CorrespondentFolders == nil {
		config.CorrespondentFolders = map[string]string{}
	}

	return SaveFullConfig(config)
}
