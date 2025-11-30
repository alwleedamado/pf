package cmd

import (
	"os"
	"path/filepath"
)

func cleanDirectory(path string) error {
	content, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range content {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			if err := os.RemoveAll(fullPath); err != nil {
				return err
			}
		} else {
			if err := os.Remove(fullPath); err != nil {
				return err
			}
		}
	}
	return nil
}
