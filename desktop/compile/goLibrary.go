package compile

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type GoLibrary struct {
	Path    string
	Library *Library

	file *os.File
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

	var libData Library
	err = yaml.Unmarshal(data, &libData)
	if err != nil {
		return err
	}

	gl.Library = &libData
	return nil
}

func (gl *GoLibrary) Dispose() {
	gl.file.Close()
	gl = nil
}
