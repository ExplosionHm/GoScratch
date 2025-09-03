package main

import (
	"embed"

	_log "log"
	"opticode/desktop"
	"opticode/desktop/log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := desktop.NewApp(log.NewLogger())
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
			app.Tree,
		},
		OnStartup: app.OnStartup,
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})
	if err != nil {
		_log.Fatal(err)
	}
}
