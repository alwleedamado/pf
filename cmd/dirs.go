package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/alwleedamado/pf/internal"
	"github.com/spf13/cobra"
)

var dirs = internal.GetDirectories()

func GetDirUsage() (map[string]string, error) {
	sizes := make(map[string]int64, len(dirs))
	for _, dir := range dirs {
		expanded := internal.ExpandTilde(dir)

		err := filepath.WalkDir(expanded, func(path string, d fs.DirEntry, err error) error {
			if os.IsNotExist(err) {
				return nil
			}
			if err != nil {
				return err
			}

			if d.Type().IsRegular() {
				info, err := d.Info()
				if err != nil {
					if os.IsNotExist(err) {
						return nil
					}
					return err
				}
				sizes[expanded] += info.Size()
			}
			return nil
		})
		if err != nil {
			delete(sizes, expanded)
		}
	}
	ret := make(map[string]string)
	for i, v := range sizes {
		ret[i] = internal.HumanBytes(v)
	}
	return ret, nil
}

func CleanDirectory(path string) error {
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

var listDirs bool
var cleanDirs bool
var targetDir string
var DirCommand = &cobra.Command{
	Use:   "dirs",
	Short: "manage directories",
	Long:  "manage cache and build directories",
	Run: func(cmd *cobra.Command, args []string) {
		if listDirs {
			usages, _ := GetDirUsage()
			for i, v := range usages {
				fmt.Printf("%-*s %-*v\n",20,i,5,v)
			}
			return
		}
	},
}

func InitDirCommand() {
	DirCommand.Flags().BoolVarP(&listDirs, "list", "l", false, "list directories")
	DirCommand.Flags().StringVarP(&targetDir, "clean", "c", "", "clean a specific cache directory")
}
