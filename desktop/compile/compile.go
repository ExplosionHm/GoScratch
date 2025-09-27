package compile

import (
	"fmt"
	"log"
	"opticode/desktop/tree"
)

type ErrorCode uint8

const (
	Warn ErrorCode = iota
	Fatal
)

type Error struct {
	err  string
	Code ErrorCode
}

func Err(c ErrorCode, format string, v ...any) *Error {
	return &Error{
		err:  fmt.Sprintf(format, v...),
		Code: c,
	}
}

func (e *Error) ErrorLevel() string {
	switch e.Code {
	case Warn:
		return "WARN"
	case Fatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func (e *Error) Error() string {
	return e.ErrorLevel() + " | " + e.err
}

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

	return out, nil
}

func nearestArgument(length int, i int) int {
	if i <= 0 || length < i {
		// is within length
		return i
	}

	return length - 1 // Return last
}
