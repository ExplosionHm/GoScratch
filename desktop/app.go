package desktop

import (
	"context"
	"log"
	"opticode/desktop/compile"
	"opticode/desktop/compile/golang"
	"opticode/desktop/project"
	"opticode/desktop/tree"
	"opticode/utils"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

type App struct {
	ctx     context.Context
	Tree    *tree.Tree
	Project *project.Project
}

func NewApp() *App {
	t := tree.NewTree()

	return &App{
		Tree: t,
	}
}

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenProject(dir string) (*project.Project, error) {
	//! Implement
	return project.NewProject(project.ProjectHeader{
		Id:        "test",
		Name:      "TEST",
		Libraries: map[string]project.LibraryPathFlag{},
	}, ""), nil
}

func (a *App) GenerateCode(dir string) error {
	log.Println("Enter gen")
	start := time.Now()
	out, err := compile.Run(a.Tree, compile.Golang)
	if err != nil {
		return err
	}
	end := time.Since(start).Milliseconds()
	log.Println("Server sent: "+out, "Completed in "+strconv.Itoa(int(end))+"ms")

	if !utils.FileExists(path.Join(dir, "go.mod")) {
		err = golang.InitGoProject(dir, "placeholder")
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(path.Join(dir, "main.go"), []byte(out), os.ModeAppend)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "run", ".")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
