package desktop

import (
	"context"
	"log"
	"opticode/desktop/compile"
	"opticode/desktop/tree"
	"opticode/utils"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

type Project struct {
}

type App struct {
	ctx     context.Context
	Tree    *tree.Tree
	Project Project
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

type Response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error,omitempty"`
}

func (a *App) GenerateCode(outdir string) error {
	start := time.Now()
	out, err := compile.Run(a.Tree, compile.Golang)
	if err != nil {
		return err
	}
	end := time.Since(start).Nanoseconds()
	log.Println("Server sent: "+out, "Completed in "+strconv.Itoa(int(end))+"ns")

	if !utils.FileExists(path.Join(outdir, "go.mod")) {
		err = compile.InitGoProject(outdir, "placeholder")
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(path.Join(outdir, "main.go"), []byte(out), os.ModeAppend)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "run", ".")
	cmd.Dir = outdir
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
