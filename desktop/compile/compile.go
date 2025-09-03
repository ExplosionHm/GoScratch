package compile

import (
	"fmt"
	"opticode/desktop/tree"
)

const StandardLibaryPath = ""

type Language string

const (
	Golang = "Go"
)

// Code generation pipeline:
// Validate -> Translate -> Format
func Run(tree *tree.Tree, lang Language) (string, error) {
	var out string
	var err error

	// Translate
	switch lang {
	case Golang:
		// Validate
		//! Implement

		standard := NewGoLibary("libraries/go_standard.yml")
		err = standard.Open()
		if err != nil {
			return "", err
		}

		gen := NewGoGenerator(tree, []*GoLibrary{standard}, "	")
		out, err = gen.Generate()
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("language not supported: %s", lang)
	}

	// Format
	//! Implement

	return out, nil
}
