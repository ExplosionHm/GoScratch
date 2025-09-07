package compile

import (
	"fmt"
	"log"
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
		log.Println("Enter golang")
		standard := NewGoLibary("libraries/go_standard.json")
		err = standard.Open()
		if err != nil {
			return "", err
		}

		defs := NewGoLibary("C:\\Users\\explo\\OneDrive\\Documents\\project\\definitions.json")
		err = defs.Open()
		if err != nil {
			return "", err
		}

		log.Println("Passed lib")

		gen := NewGoGenerator(tree, []*GoLibrary{standard, defs}, "	")
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

func nearestArgument(length int, i int) int {
	if i <= 0 || length < i {
		// is within length
		return i
	}

	return length - 1 // Return last
}
