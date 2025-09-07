package main

import (
	"embed"
	"os"

	"log"
	"opticode/desktop"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Debugging enabled")
	// This doesn't work that well
	/* file, err := os.OpenFile("./logs/"+time.Now().Format(time.DateOnly)+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file) */

	app := desktop.NewApp()
	err := wails.Run(&options.App{
		Title:            "Opticode",
		Width:            1024,
		Height:           768,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			app,
			app.Tree,
		},
		OnStartup: app.OnStartup,
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
