package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/guessi/kubectl-grep/pkg/options"
)

var (
	output                  string
	rootCmdDescriptionShort = "Filter Kubernetes resources by matching their names"
	rootCmdDescriptionLong  = `Filter Kubernetes resources by matching their names

More info: https://github.com/guessi/kubectl-grep
`

	rootCmdExamples = `
List all pods in default namespace
$ kubectl grep pods

List all pods in all namespaces
$ kubectl grep pods -A

List all pods in namespace "start-lab" which contains keyword "flash"
$ kubectl grep pods -n star-lab flash
`
)

// generic search options handler
var searchOptions = options.NewSearchOptions()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "kubectl-grep",
	Short:   rootCmdDescriptionShort,
	Long:    rootCmdDescriptionLong,
	Example: rootCmdExamples,
}

func init() {
	// Global Flags
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.Namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	rootCmd.PersistentFlags().BoolVarP(
		&searchOptions.AllNamespaces, "all-namespaces", "A", false,
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
