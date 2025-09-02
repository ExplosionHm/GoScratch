package desktop

import (
	"bytes"
	"context"
	"net/url"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

type App struct {
	ctx context.Context
	Log logger.Logger
}

func NewApp(l logger.Logger) *App {
	return &App{
		Log: l,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OnStartup(ctx context.Context) {
	a.Log.Info("I have started up")
}

// ! Implement
func (a *App) TranspileFile(*url.URL) *bytes.Buffer {
	a.Log.Debug("Not implemented")
	return nil
}

// ! Implement
func (a *App) Transpile(buf *bytes.Buffer) *bytes.Buffer {
	a.Log.Debug("Not implemented")
	return nil
}

// ! Implement
// Returns executable path
func (a *App) Compile(buf *bytes.Buffer) *string {
	a.Log.Debug("Not implemented")
	return nil
}
