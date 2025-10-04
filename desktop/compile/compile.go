package compile

import (
	"fmt"
	"log"
	"opticode/desktop/compile/golang"
	"opticode/desktop/tree"
)

type Language string

const (
	Golang Language = "Go"
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
		log.Println("Enter golang")
		standard := golang.NewGoLibary("libraries/go_standard.json")
		err = standard.Open()
		if err != nil {
			return "", err
		}

		defs := golang.NewGoLibary("C:\\Users\\explo\\OneDrive\\Documents\\project\\definitions.json")
		err = defs.Open()
		if err != nil {
			return "", err
		}
		log.Println("Passed lib")

		gen := golang.NewGoGenerator(tree, []*golang.GoLibrary{standard, defs}, "	")
		out, err = gen.Generate()
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("language not supported: %s", lang)
	}

	return out, nil
}
