package cmd

import (
	"github.com/guessi/kubectl-search/utils/search"

	"github.com/spf13/cobra"
)

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Search Pods by keyword, by namespace",
	Run: func(cmd *cobra.Command, args []string) {
		searchPods(args)
	},
}

func init() {
	rootCmd.AddCommand(podsCmd)
	podsCmd.Flags().StringVarP(
		&namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	podsCmd.Flags().StringVarP(
		&selector, "selector", "l", "",
		"Selector (label query) to filter on. (e.g. -l key1=value1,key2=value2)")
	podsCmd.Flags().StringVar(
		&fieldSelector, "field-selector", "",
		"Selector (field query) to filter on. (e.g. --field-selector key1=value1,key2=value2)")
	podsCmd.Flags().BoolVar(
		&allNamespaces, "all-namespaces", false,
		"If present, list the requested object(s) across all namespaces.")
}

func searchPods(args []string) {
	var searchKeyword string

	if len(args) >= 1 && args[0] != "" {
		searchKeyword = trimQuoteAndSpace(args[0])
	}

	search.Pods(namespace, allNamespaces, selector, fieldSelector, searchKeyword)
}
