package cmd

import (
	"github.com/guessi/kubectl-search/pkg/search"

	"github.com/spf13/cobra"
)

// daemonsetsCmd represents the pods command
var daemonsetsCmd = &cobra.Command{
	Use:   "daemonsets",
	Short: "Search Daemonsets by keyword, by namespace",
	Run: func(cmd *cobra.Command, args []string) {
		searchDaemonsets(args)
	},
}

func init() {
	rootCmd.AddCommand(daemonsetsCmd)
	daemonsetsCmd.Flags().StringVarP(
		&output, "output", "o", "",
		"Output format.")
}

func searchDaemonsets(args []string) {
	var searchKeyword string

	if len(args) >= 1 && args[0] != "" {
		searchKeyword = trimQuoteAndSpace(args[0])
	}

	search.Daemonsets(namespace, allNamespaces, selector, fieldSelector, searchKeyword, output == "wide")
}
