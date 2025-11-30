package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/alwleedamado/pf/internal"
	"github.com/alwleedamado/pf/internal/provider"
	"github.com/spf13/cobra"
)

func GetDirUsage() []provider.Usage {
	gp := provider.GlobalProviders()
	var usages []provider.Usage
	for _, p := range gp {
		for _, i := range p.Paths() {
			usage, _ := provider.Scan(i, p.Name())
			usages = append(usages, usage)
		}
	}
	return usages
}

var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "List global & project caches",
	Run: func(cmd *cobra.Command, args []string) {

		usages := GetDirUsage()
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		defer w.Flush()
		fmt.Fprintln(w, "Type\tProvider\tSize\tPath")
		for _, v := range usages {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", "Global", v.Label, internal.HumanBytes(v.Size), v.Path)
		}
	},
}

func InitListCmd() {
}
