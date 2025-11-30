package cmd

import (
	"os"
	"path/filepath"

	"github.com/alwleedamado/pf/internal/provider"
	"github.com/spf13/cobra"
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

var dryRun bool
var CleanCommand = &cobra.Command{
	Use:   "clean",
	Short: "clean directoris",
	Run: func(cmd *cobra.Command, args []string) {
		if !dryRun {
			for _, p := range provider.GlobalProviders() {
				for _, path := range p.Paths() {
					cleanDirectory(path)
				}
			}
		}
	},
}

func InitCleanCmd() {
	CleanCommand.Flags().BoolVar(&dryRun, "dry-run", false, "don't clean anything")
}
