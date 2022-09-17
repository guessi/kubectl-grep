package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

var (
	gitVersion   = "v0.0.0"
	goVersion    = "v0.0.0"
	buildTime    = "undefined"
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
	r, _ := regexp.Compile(`v[0-9]\.[0-9]+\.[0-9]+`)
	versionInfo := r.FindString(gitVersion)
	if shortVersion {
		fmt.Println(versionInfo)
	} else {
		fmt.Println("kubectl-grep", versionInfo)
		fmt.Println(" Git Commit:", gitVersion)
		fmt.Println(" Build with:", goVersion)
		fmt.Println(" Build time:", buildTime)
	}
}
