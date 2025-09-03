package desktop

import (
	"context"
	"opticode/desktop/compile"
	"opticode/desktop/log"
	"opticode/desktop/tree"
	"os"
	"os/exec"
	"path"
)

type Project struct {
}

type App struct {
	ctx     context.Context
	Log     *log.Logger
	Tree    *tree.Tree
	Project Project
}

func NewApp(l *log.Logger) *App {
	t := tree.NewTree(l)

	return &App{
		Log:  l,
		Tree: t,
	}
}

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
	a.Log.Info("STARTUP")

	if len(a.Tree.Nodes) == 0 {
		a.Log.Debug("tree is empty")
	}
}

type Response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error,omitempty"`
}

func (a *App) GenerateCode(outdir string) error {
	out, err := compile.Run(a.Tree, compile.Golang)
	if err != nil {
		return err
	}
	a.Log.Print("Server sent: " + out)

	err = compile.InitGoProject(outdir, "placeholder")
	if err != nil {
		return err
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
