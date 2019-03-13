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
}

func searchHpas(args []string) {
	var searchKeyword string

	if len(args) >= 1 && args[0] != "" {
		searchKeyword = trimQuoteAndSpace(args[0])
	}

	search.Hpas(namespace, allNamespaces, selector, fieldSelector, searchKeyword)
}
