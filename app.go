package main

// Add build tags

import (
	"context"
	"log"
	"opticode/compile/golang"
	"opticode/project"
	"time"
)

type App struct {
	ctx     context.Context
	Project *project.Project
}

// All methods of App are exposed to the frontend

func NewApp() *App {
	return &App{}
}

func (a *App) Initalize(ctx context.Context) {
	a.ctx = ctx
	a.Project = project.NewProject(project.Header{}, project.Appearances{}, "")
}

func (a *App) Exit(ctx context.Context) {
	log.Println("Shutting down...")
}

func (a *App) OpenProject(projectFile string) (*project.Project, error) {
	return project.LoadProjectFromFile(projectFile)
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
