package src

import (
	"net/http"
	"os"
	"path/filepath"
)

// StartPDFServer startet einen dedizierten lokalen Webserver nur für die PDFs
func StartPDFServer(workDirProvider func() string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/pdf/", func(rw http.ResponseWriter, req *http.Request) {
		// CORS-Header erlauben, damit das Wails-Frontend die Datei lesen darf
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		filename := req.URL.Path[5:] // Schneidet "/pdf/" ab
		workDir := workDirProvider()

		if workDir == "" || filename == "" {
			http.Error(rw, "Pfad oder Dateiname fehlt", http.StatusBadRequest)
			return
		}

		filePath := filepath.Clean(filepath.Join(workDir, "900-Eingang", filename))
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
		_ = http.ListenAndServe("127.0.0.1:34999", mux)
	}()
}
