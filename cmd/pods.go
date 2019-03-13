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
		&output, "output", "o", "",
		"Output format.")
}

func searchPods(args []string) {
	var searchKeyword string

	if len(args) >= 1 && args[0] != "" {
		searchKeyword = trimQuoteAndSpace(args[0])
	}

	search.Pods(namespace, allNamespaces, selector, fieldSelector, searchKeyword, output == "wide")
}
