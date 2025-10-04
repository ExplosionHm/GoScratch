package golang

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"encoding/json"
)

type FuncDef struct {
	Added     float64    `yaml:"added"`
	Types     []TypeDef  `yaml:"types,omitempty"`
	Arguments [][]string `yaml:"arguments,omitempty"`
	Returns   *string    `yaml:"returns"`
}

type TypeDef map[string]string

type GoLibrary struct {
	Id      string             `json:"id"`
	Version string             `json:"version"`
	Types   []string           `json:"types"`
	Funcs   map[string]FuncDef `json:"funcs"`

	file *os.File `json:"-"`
	Path string   `json:"-"`
}

func NewGoLibary(path string) *GoLibrary {
	return &GoLibrary{
		Path: path,
	}
}

func (gl *GoLibrary) Open() error {
	var err error
	gl.file, err = os.OpenFile(gl.Path, 0, os.ModeAppend)
	if err != nil {
		return err

	}
	data, err := io.ReadAll(gl.file)
	if err != nil {
		return err
	}

	var libData GoLibrary
	err = json.Unmarshal(data, &libData)
	if err != nil {
		return err
	}

	(*gl) = libData
	return nil
}

func (l *GoLibrary) LookupFunc(opid string) *FuncDef {
	if def, ok := l.Funcs[opid]; ok {
		return &def
	}
	return nil
}

func (gl *GoLibrary) Dispose() {
	gl.file.Close()
	gl = nil
}

func InitGoProject(dir string, name string) error {
	info, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("cannot access dir %q: %w", dir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("path %q is not a directory", dir)
	}

	if name == "" {
		return fmt.Errorf("module name cannot be empty")
	}

	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = dir

	std, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod init failed: %w\nOutput:\n%s", err, string(std))
	}
	println("std out:", string(std))

	mainPath := path.Join(dir, "main.go")
	if _, err := os.Stat(mainPath); err == nil {
		return fmt.Errorf("main.go already exists at %q", mainPath)
	}

	f, err := os.Create(mainPath)
	if err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}
	defer f.Close()

	return nil
}
