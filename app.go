package main

// Add build tags

import (
	"context"
	"encoding/base64"
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

func (a *App) Compile(data string, dir string) (uint8, error) {
	buf, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return 2, err
	}

	start := time.Now()
	files, err := golang.Compile(&buf)
	if err != nil {
		return 4, err
	}
	log.Printf("Time elapsed: %dms", time.Since(start).Milliseconds())
	log.Printf("# of files: %d", len(files))

	return 0, nil
}
