package cmd

import (
	"github.com/guessi/kubectl-grep/pkg/resources"
	"github.com/guessi/kubectl-grep/pkg/utils"

	"github.com/spf13/cobra"
)

var (
	// daemonsetsCmd represents the pods command
	daemonsetsCmd = &cobra.Command{
		Use:     "daemonsets",
		Aliases: []string{"ds", "daemonset"},
		Short:   "Search Daemonsets by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "daemonsets")
		},
	}
	// deploymentsCmd represents the pods command
	deploymentsCmd = &cobra.Command{
		Use:     "deployments",
		Aliases: []string{"deploy", "deployment"},
		Short:   "Search Deployments by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "deployments")
		},
	}
	// hpasCmd represents the hpas command
	hpasCmd = &cobra.Command{
		Use:     "hpas",
		Aliases: []string{"hpa"},
		Short:   "Search HPAs by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "hpas")
		},
	}
	// nodesCmd represents the nodes command
	nodesCmd = &cobra.Command{
		Use:     "nodes",
		Aliases: []string{"no", "nodes"},
		Short:   "Search Nodes by keyword",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "nodes")
		},
	}
	// podsCmd represents the pods command
	podsCmd = &cobra.Command{
		Use:     "pods",
		Aliases: []string{"po", "pod"},
		Short:   "Search Pods by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "pods")
		},
	}
	// configmapsCmd represents the configmaps command
	configmapsCmd = &cobra.Command{
		Use:     "configmaps",
		Aliases: []string{"cm", "configmap"},
		Short:   "Search ConfigMaps by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "configmaps")
		},
	}
	// statefulsetsCmd represents the statefulsets command
	statefulsetsCmd = &cobra.Command{
		Use:     "statefulsets",
		Aliases: []string{"sts", "statefulset"},
		Short:   "Search Statefulsets by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "statefulsets")
		},
	}
)

func init() {
	rootCmd.AddCommand(daemonsetsCmd)
	rootCmd.AddCommand(deploymentsCmd)
	rootCmd.AddCommand(hpasCmd)
	rootCmd.AddCommand(nodesCmd)
	rootCmd.AddCommand(podsCmd)
	rootCmd.AddCommand(configmapsCmd)
	rootCmd.AddCommand(statefulsetsCmd)

	daemonsetsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")
	deploymentsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")
	nodesCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")
	podsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")
	statefulsetsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")
}

func resourceSearch(args []string, resourceType string) {
	var keyword string

	if len(args) >= 1 && args[0] != "" {
		keyword = utils.TrimQuoteAndSpace(args[0])
	}

	switch resourceType {
	case "daemonsets":
		resources.Daemonsets(searchOptions, keyword, output == "wide")
	case "deployments":
		resources.Deployments(searchOptions, keyword, output == "wide")
	case "hpas":
		resources.Hpas(searchOptions, keyword)
	case "nodes":
		resources.Nodes(searchOptions, keyword, output == "wide")
	case "pods":
		resources.Pods(searchOptions, keyword, output == "wide")
	case "configmaps":
		resources.ConfigMaps(searchOptions, keyword)
	case "statefulsets":
		resources.Statefulsets(searchOptions, keyword, output == "wide")
	default:
		break
	}
}
