package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/guessi/kubectl-grep/pkg/options"
)

var (
	cfgFile            string
	output             string
	rootCmdDescription = `kubectl plugins for searching Kubernetes resources

Find more information at: https://github.com/guessi/kubectl-grep
	`
)

// generic search options handler
var searchOptions = options.NewSearchOptions()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-grep",
	Short: "kubectl plugins for searching Kubernetes resources",
	Long:  rootCmdDescription,
}

func init() {
	// Global Flags
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.Namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	rootCmd.PersistentFlags().BoolVar(
		&searchOptions.AllNamespaces, "all-namespaces", false,
		"If present, list the requested object(s) across all namespaces.")
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.Selector, "selector", "l", "",
		"Selector (label query) to filter on. (e.g. -l key1=value1,key2=value2)")
	rootCmd.PersistentFlags().StringVar(
		&searchOptions.FieldSelector, "field-selector", "",
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
