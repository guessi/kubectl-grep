package cmd

import (
	"github.com/guessi/kubectl-search/pkg/search"

	"github.com/spf13/cobra"
)

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Search Nodes by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		searchNodes(args)
	},
}

func init() {
	rootCmd.AddCommand(nodesCmd)
	nodesCmd.Flags().StringVarP(
		&output, "output", "o", "",
		"Output format.")
}

func searchNodes(args []string) {
	var searchKeyword string

	if len(args) >= 1 && args[0] != "" {
		searchKeyword = trimQuoteAndSpace(args[0])
	}

	search.Nodes(selector, fieldSelector, searchKeyword, output == "wide")
}
