package fsHandler

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
)

func MkDir(path string, name string) (dir string, crError error) {
	// Create the directory
	newPath := filepath.Join(path, name)
	err := os.MkdirAll(newPath, 0755)
	if err != nil {
		log.Printf("Error creating directory: %s", err)
		return "" ,fmt.Errorf("failed to create directory: %w", err)
	}
	return newPath, nil
}