package desktop

import (
	"context"
	"opticode/desktop/project"
	"os"
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

func (a *App) GenerateCode(dir string) error {
	/* log.Println("Enter gen")
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
	*/
	return nil
}
