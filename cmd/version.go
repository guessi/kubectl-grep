package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appVersion   = "1.0.0"
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
		fmt.Println(appVersion)
	} else {
		fmt.Println("kubectl-search version:", appVersion)
	}
}
