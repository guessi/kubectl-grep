package cmd

import (
	"github.com/spf13/cobra"

	"github.com/guessi/kubectl-search/utils/search"
)

// deploymentsCmd represents the pods command
var deploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "Search Deployments by keyword, by namespace",
	Run: func(cmd *cobra.Command, args []string) {
		searchDeployments(args)
	},
}

func init() {
	rootCmd.AddCommand(deploymentsCmd)
	deploymentsCmd.Flags().StringVarP(
		&namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	deploymentsCmd.Flags().StringVarP(
		&selector, "selector", "l", "",
		"Selector (label query) to filter on. (e.g. -l key1=value1,key2=value2)")
	deploymentsCmd.Flags().StringVar(
		&fieldSelector, "field-selector", "",
		"Selector (field query) to filter on. (e.g. --field-selector key1=value1,key2=value2)")
	deploymentsCmd.Flags().BoolVar(
		&allNamespaces, "all-namespaces", false,
		"If present, list the requested object(s) across all namespaces.")
}

func searchDeployments(args []string) {
	var searchKeyword string

	if len(args) >= 1 && args[0] != "" {
		searchKeyword = trimQuoteAndSpace(args[0])
	}

	search.Deployments(namespace, allNamespaces, selector, fieldSelector, searchKeyword)
}
