package src

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// GetOldestInboxPDF sucht die älteste PDF-Datei im Ordner 900-Eingang
func GetOldestInboxPDF(workDir string) (string, error) {
	inboxPath := filepath.Join(workDir, "900-Eingang")

	files, err := os.ReadDir(inboxPath)
	if err != nil {
		return "", err
	}

	type fileInfo struct {
		name string
		modTime os.FileInfo
	}

	var pdfs []os.FileInfo

	// Wir holen uns die echten Datei-Infos für den Zeitstempel
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".pdf") {
			info, err := file.Info()
			if err == nil {
				pdfs = append(pdfs, info)
			}
		}
	}

	// Wenn keine PDFs da sind, geben wir einen leeren String zurück
	if len(pdfs) == 0 {
		return "", nil
	}

	// Sortieren: Älteste Datei (kleinstes ModTime) zuerst
	sort.Slice(pdfs, func(i, j int) bool {
		return pdfs[i].ModTime().Before(pdfs[j].ModTime())
	})

	// Rückgabe des ältesten Dateinamens
	return pdfs[0].Name(), nil
}
