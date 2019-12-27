package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// GitCommit refers to commit hash at the moment of build
var GitCommit string

// Version of Mysh
var Version string

// BuildTime of Mysh
var BuildTime string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Mysh",
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 14, 1, 1, ' ', 0)
		fmt.Fprintf(w, "Version:\t%s\n", Version)
		fmt.Fprintf(w, "GitCommit:\t%s\n", GitCommit)
		fmt.Fprintf(w, "Built:\t%s\n", BuildTime)
		w.Flush()
	},
}
