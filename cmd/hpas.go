package cmd

import (
	"github.com/guessi/kubectl-search/utils/search"

	"github.com/spf13/cobra"
)

// hpasCmd represents the hpas command
var hpasCmd = &cobra.Command{
	Use:   "hpas",
	Short: "Search HPAs by keyword, by namespace",
	Run: func(cmd *cobra.Command, args []string) {
		searchHpas(args)
	},
}

func init() {
	rootCmd.AddCommand(hpasCmd)
	hpasCmd.Flags().StringVarP(
		&namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	hpasCmd.Flags().StringVarP(
		&selector, "selector", "l", "",
		"Selector (label query) to filter on. (e.g. -l key1=value1,key2=value2)")
	hpasCmd.Flags().StringVar(
		&fieldSelector, "field-selector", "",
		"Selector (field query) to filter on. (e.g. --field-selector key1=value1,key2=value2)")
	hpasCmd.Flags().BoolVar(
		&allNamespaces, "all-namespaces", false,
		"If present, list the requested object(s) across all namespaces.")
}

func searchHpas(args []string) {
	var searchKeyword string

	if len(args) >= 1 && args[0] != "" {
		searchKeyword = trimQuoteAndSpace(args[0])
	}

	search.Hpas(namespace, allNamespaces, selector, fieldSelector, searchKeyword)
}
