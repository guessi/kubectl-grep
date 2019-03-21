package cmd

import (
	"github.com/guessi/kubectl-search/pkg/search"

	"github.com/spf13/cobra"
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
		&output, "output", "o", "",
		"Output format.")
}

func searchDeployments(args []string) {
	var searchKeyword string

	if len(args) >= 1 && args[0] != "" {
		searchKeyword = trimQuoteAndSpace(args[0])
	}

	search.Deployments(namespace, allNamespaces, selector, fieldSelector, searchKeyword, output == "wide")
}
