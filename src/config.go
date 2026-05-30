package src

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// AppConfig hält die session-übergreifenden Einstellungen
type AppConfig struct {
	WorkDir              string            `json:"workDir"`
	Correspondents       []string          `json:"correspondents"`
	CorrespondentFolders map[string]string `json:"correspondentFolders"`
}

// GetConfigPath ermittelt den plattformspezifischen Konfigurationspfad:
//
//	macOS:   ~/Library/Application Support/eu.lowsky.sort-pdf/config.json
//	Windows: %APPDATA%/eu.lowsky.sort-pdf/config.json
//	Linux:   ~/.config/eu.lowsky.sort-pdf/config.json
func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(configDir, "eu.lowsky.sort-pdf")
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

	bytes, err := os.ReadFile(configPath)
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

	return os.WriteFile(configPath, bytes, 0644)
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
		// Config existiert noch nicht (erster Start) – mit leeren Werten beginnen
		config = AppConfig{
			Correspondents:       []string{},
			CorrespondentFolders: map[string]string{},
		}
	}

	config.WorkDir = path

	return SaveFullConfig(config)
}
