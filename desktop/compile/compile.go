package compile

import (
	"bytes"
	"fmt"
	"opticode/desktop/compile/golang"
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

		golang.Compile(nil, programBuf.Bytes())
	default:
		return "", fmt.Errorf("language not supported: %s", lang)
	}

	return out, nil
}
