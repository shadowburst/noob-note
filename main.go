package main

import (
	"context"
	"embed"
	"log"
	"noob-note/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the backend structure
	backend := backend.NewBackend()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "noob-note",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			err := backend.Start()
			if err != nil {
				log.Fatal(err)
			}
		},
		OnShutdown: func(ctx context.Context) {
			err := backend.Stop()
			if err != nil {
				log.Fatal(err)
			}
		},
		Bind: []interface{}{
			backend,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
