package src

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// StartPDFServer startet einen dedizierten lokalen Webserver nur für die PDFs
func StartPDFServer(workDirProvider func() string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/pdf/", func(rw http.ResponseWriter, req *http.Request) {
		// CORS-Header erlauben, damit das Wails-Frontend die Datei lesen darf
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		rawFilename := req.URL.Path[5:] // Schneidet "/pdf/" ab
		filename, err := url.PathUnescape(rawFilename)
		if err != nil || filename == "" {
			http.Error(rw, "Ungültiger Dateiname", http.StatusBadRequest)
			return
		}
		workDir := workDirProvider()

		if workDir == "" {
			http.Error(rw, "Pfad fehlt", http.StatusBadRequest)
			return
		}

		inboxPath, err := filepath.EvalSymlinks(filepath.Join(workDir, "900-Eingang"))
		if err != nil {
			http.Error(rw, "Pfadfehler", http.StatusInternalServerError)
			return
		}

		filePath, err := filepath.EvalSymlinks(filepath.Join(inboxPath, filename))
		if err != nil {
			http.Error(rw, "Datei nicht gefunden", http.StatusNotFound)
			return
		}

		// Path-Traversal-Schutz: Nur Dateien innerhalb von 900-Eingang ausliefern
		if !strings.HasPrefix(filePath, inboxPath) {
			http.Error(rw, "Zugriff verweigert", http.StatusForbidden)
			return
		}

		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(rw, "Datei nicht gefunden", http.StatusNotFound)
			return
		}

		rw.Header().Set("Content-Type", "application/pdf")
		rw.Write(fileBytes)
	})

	// Startet den Server im Hintergrund auf Port 34999
	go func() {
		if err := http.ListenAndServe("127.0.0.1:34999", mux); err != nil {
			log.Printf("PDF-Server Fehler: %v", err)
		}
	}()
}
