package main

import (
	"context"
	"embed"
	"sort-pdf/src" // Importiert unseren neuen src-Ordner

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Erstelle eine Instanz unserer neuen fokussierten Brücke
	bridge := src.NewWorkspaceBridge()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sort-pdf",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			// Setzt den Wails-Context in unserer Brücke, damit wir Dialoge öffnen können
			bridge.SetContext(ctx)
		},
		Bind: []interface{}{
			bridge, // Bindet die Brücke an das Frontend an
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
