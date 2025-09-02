package main

import (
	"embed"
	"log"

	"opticode/desktop"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := desktop.NewApp(logger.NewDefaultLogger())
	err := wails.Run(&options.App{
		Title:            "Opticode",
		Width:            1024,
		Height:           768,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Logger: app.Log,
		Bind: []interface{}{
			app,
		},
		OnStartup: app.OnStartup,
	})
	if err != nil {
		log.Fatal(err)
	}

	app.Log.Info("Test")
}
