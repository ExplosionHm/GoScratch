package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		log.Printf("failed to verify file %s: %v\n", filePath, err)
		return false
	}
	return true
}

func Nearest(length int, i int) int {
	if i <= 0 || length < i {
		// is within length
		return i
	}

	return length - 1 // Return last
}

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
