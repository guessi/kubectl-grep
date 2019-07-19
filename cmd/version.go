package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version      = "undefined"
	shortVersion bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number",
	Run: func(cmd *cobra.Command, args []string) {
		showVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&shortVersion, "short", "s", false, "short version output")
}

func showVersion() {
	if shortVersion {
		fmt.Println(version)
	} else {
		fmt.Println("kubectl-grep version:", version)
	}
}
