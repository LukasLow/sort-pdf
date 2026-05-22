package main

import (
	"context"
	"embed"
	"sort-pdf/src"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	bridge := src.NewWorkspaceBridge()

	err := wails.Run(&options.App{
		Title:  "sort-pdf",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets, // Zurück zum Standard-Verhalten von Wails
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			bridge.SetContext(ctx)

			// Startet unseren PDF-Server und übergibt eine Funktion,
			// die immer das aktuell ausgewählte Arbeitsverzeichnis liefert
			src.StartPDFServer(func() string {
				return bridge.CheckInitialWorkDir()
			})
		},
		Bind: []interface{}{
			bridge,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
