package src

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type pdfEntry struct {
	relPath string
	modTime time.Time
}

// GetOldestInboxPDF sucht die älteste PDF-Datei rekursiv im Ordner 900-Eingang
// inklusive aller Unterordner. Der Rückgabewert ist ein relativer Pfad
// (z.B. "unterordner/datei.pdf").
func GetOldestInboxPDF(workDir string) (string, error) {
	inboxPath := filepath.Join(workDir, "900-Eingang")

	if _, err := os.Stat(inboxPath); os.IsNotExist(err) {
		return "", nil
	}

	var pdfs []pdfEntry

	err := filepath.WalkDir(inboxPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(strings.ToLower(d.Name()), ".pdf") {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return nil
		}
		relPath, err := filepath.Rel(inboxPath, path)
		if err != nil {
			return nil
		}
		pdfs = append(pdfs, pdfEntry{relPath, info.ModTime()})
		return nil
	})

	if err != nil {
		return "", err
	}
	if len(pdfs) == 0 {
		return "", nil
	}

	sort.Slice(pdfs, func(i, j int) bool {
		return pdfs[i].modTime.Before(pdfs[j].modTime)
	})

	return pdfs[0].relPath, nil
}
