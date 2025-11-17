package main

import (
	"embed"
	"os"

	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	log.SetOutput(os.Stdout) // Set wails logging to default std output

	app := NewApp()
	err := wails.Run(&options.App{
		Title:            "Opticode",
		Width:            1024,
		Height:           768,
		MinWidth:         480,
		MinHeight:        360,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []any{
			app,
			//app.Project,
		},
		OnStartup:  app.Initalize,
		OnShutdown: app.Exit,
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
