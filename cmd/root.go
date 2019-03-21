package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile            string
	allNamespaces      bool
	namespace          string
	selector           string
	fieldSelector      string
	output             string
	rootCmdDescription = `kubectl plugins for searching Kubernetes resources

Find more information at: https://github.com/guessi/kubectl-search
	`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-search",
	Short: "kubectl plugins for searching Kubernetes resources",
	Long:  rootCmdDescription,
}

func init() {
	// Global Flags
	rootCmd.PersistentFlags().StringVarP(
		&namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	rootCmd.PersistentFlags().BoolVar(
		&allNamespaces, "all-namespaces", false,
		"If present, list the requested object(s) across all namespaces.")
	rootCmd.PersistentFlags().StringVarP(
		&selector, "selector", "l", "",
		"Selector (label query) to filter on. (e.g. -l key1=value1,key2=value2)")
	rootCmd.PersistentFlags().StringVar(
		&fieldSelector, "field-selector", "",
		"Selector (field query) to filter on. (e.g. --field-selector key1=value1,key2=value2)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
