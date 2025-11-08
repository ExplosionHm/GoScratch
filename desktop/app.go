package desktop

import (
	"context"
	"log"
	"opticode/desktop/compile/golang"
	"opticode/desktop/project"
	"os"
	"time"
)

type App struct {
	ctx     context.Context
	Project *project.Project
}

func NewApp() *App {
	return &App{}
}

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenProject(projectFile string) (*project.Project, error) {
	content, err := os.ReadFile(projectFile)
	if err != nil {
		return nil, err
	}
	return project.LoadProject(content), nil
}

func (a *App) Compile(buf []byte, lut map[uint32][]byte, outDir string) error {
	start := time.Now()
	files, err := golang.Compile(&buf, lut)
	if err != nil {
		return err
	}
	log.Println(len(files))

	log.Printf("Time elapsed: %dms", time.Since(start).Milliseconds())

	return nil
}
