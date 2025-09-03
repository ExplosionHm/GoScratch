package utils

import (
	"errors"
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
