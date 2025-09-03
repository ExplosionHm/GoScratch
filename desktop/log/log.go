package log

import (
	"os"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

// Print works like Sprintf.
func (l *Logger) Print(message string) {
	println(message)
}

// Trace level logging. Works like Sprintf.
func (l *Logger) Trace(message string) {
	println("TRA | " + message)
}

// Debug level logging. Works like Sprintf.
func (l *Logger) Debug(message string) {
	println("DEB | " + message)
}

// Info level logging. Works like Sprintf.
func (l *Logger) Info(message string) {
	println("INF | " + message)
}

// Warning level logging. Works like Sprintf.
func (l *Logger) Warning(message string) {
	println("WAR | " + message)
}

// Error level logging. Works like Sprintf.
func (l *Logger) Error(message string) {
	println("ERR | " + message)
}

// Fatal level logging. Works like Sprintf.
func (l *Logger) Fatal(message string) {
	println("FAT | " + message)
	os.Exit(1)
}
