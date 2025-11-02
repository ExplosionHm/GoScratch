package compile

import (
	"bytes"
	"fmt"
)

type Language string

const (
	Golang Language = "Go"
)

// Code generation pipeline:
// Validate -> Translate -> Format
func Run(programBuf *bytes.Buffer, lang Language) (string, error) {
	var out string
	//var err error

	// Translate
	switch lang {
	case Golang:
		// Validate

		/* defs := golang.NewGoLibary("C:\\Users\\explo\\OneDrive\\Documents\\project\\definitions.json")
		err = defs.Open()
		if err != nil {
			return "", err
		}
		log.Println("Passed lib")

		gen := golang.NewGoGenerator(tree, []*golang.GoLibrary{standard, defs}, "	")
		out, err = gen.Generate()
		if err != nil {
			return "", err
		} */
	default:
		return "", fmt.Errorf("language not supported: %s", lang)
	}

	return out, nil
}
