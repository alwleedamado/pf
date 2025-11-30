package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "pf",
	Short: "manage project dev files usage",
	Long:  "report or clean projects cache, build directories, and temp data",
	Run: func(command *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error occured: %s\n", err)
	}
}

func Init() {
	rootCommand.AddCommand(ListCommand)
}
