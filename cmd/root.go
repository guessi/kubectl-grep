package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile       string
	allNamespaces bool
	namespace     string
	selector      string
	fieldSelector string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-search",
	Short: "kubectl plugins for searching Kubernetes resources",
	Long: `kubectl plugins for searching Kubernetes resources

Find more information at: https://github.com/guessi/kubectl-search
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
