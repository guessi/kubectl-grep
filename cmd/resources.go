package cmd

import (
	"github.com/guessi/kubectl-grep/pkg/resources"
	"github.com/guessi/kubectl-grep/pkg/utils"

	"github.com/spf13/cobra"
)

var (
	// apps/v1
	daemonsetsCmd = &cobra.Command{
		Use:     "daemonsets",
		Aliases: []string{"ds", "daemonset"},
		Short:   "Search Daemonsets by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "daemonsets")
		},
	}
	deploymentsCmd = &cobra.Command{
		Use:     "deployments",
		Aliases: []string{"deploy", "deployment"},
		Short:   "Search Deployments by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "deployments")
		},
	}
	replicasetsCmd = &cobra.Command{
		Use:     "replicasets",
		Aliases: []string{"rs", "replicaset"},
		Short:   "Search Replicasets by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "replicasets")
		},
	}
	statefulsetsCmd = &cobra.Command{
		Use:     "statefulsets",
		Aliases: []string{"sts", "statefulset"},
		Short:   "Search Statefulsets by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "statefulsets")
		},
	}

	// autoscaling/v1
	hpasCmd = &cobra.Command{
		Use:     "hpas",
		Aliases: []string{"hpa"},
		Short:   "Search HPAs by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "hpas")
		},
	}

	// batch/v1
	jobsCmd = &cobra.Command{
		Use:     "jobs",
		Aliases: []string{"job"},
		Short:   "Search Jobs by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "jobs")
		},
	}

	// networking.k8s.io/v1
	ingressesCmd = &cobra.Command{
		Use:     "ingresses",
		Aliases: []string{"ing", "ingress"},
		Short:   "Search Ingresses by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "ingresses")
		},
	}

	// storage.k8s.io/v1
	csiDriversCmd = &cobra.Command{
		Use:     "csidrivers",
		Aliases: []string{"csidrivers"},
		Short:   "Search csidrivers by keyword",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "csidrivers")
		},
	}
	storageClassesCmd = &cobra.Command{
		Use:     "storageclasses",
		Aliases: []string{"storageclasses", "storageclasse", "sc"},
		Short:   "Search storageclasses by keyword",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "storageclasses")
		},
	}

	// v1
	configmapsCmd = &cobra.Command{
		Use:     "configmaps",
		Aliases: []string{"cm", "configmap"},
		Short:   "Search ConfigMaps by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "configmaps")
		},
	}
	nodesCmd = &cobra.Command{
		Use:     "nodes",
		Aliases: []string{"no", "nodes"},
		Short:   "Search Nodes by keyword",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "nodes")
		},
	}
	podsCmd = &cobra.Command{
		Use:     "pods",
		Aliases: []string{"po", "pod"},
		Short:   "Search Pods by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "pods")
		},
	}
	secretsCmd = &cobra.Command{
		Use:     "secrets",
		Aliases: []string{"secret"},
		Short:   "Search Secrets by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "secrets")
		},
	}
	servicesCmd = &cobra.Command{
		Use:     "services",
		Aliases: []string{"svc", "service"},
		Short:   "Search Services by keyword, by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			resourceSearch(args, "services")
		},
	}
)

func init() {
	// apps/v1
	rootCmd.AddCommand(daemonsetsCmd)
	daemonsetsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")

	rootCmd.AddCommand(deploymentsCmd)
	deploymentsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")

	rootCmd.AddCommand(replicasetsCmd)
	replicasetsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")

	rootCmd.AddCommand(statefulsetsCmd)
	statefulsetsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")

	// autoscaling/v1
	rootCmd.AddCommand(hpasCmd)

	// batch/v1
	rootCmd.AddCommand(jobsCmd)

	// networking.k8s.io/v1
	rootCmd.AddCommand(ingressesCmd)

	// storage.k8s.io/v1
	rootCmd.AddCommand(csiDriversCmd)

	rootCmd.AddCommand(storageClassesCmd)

	// v1
	rootCmd.AddCommand(configmapsCmd)

	rootCmd.AddCommand(nodesCmd)
	nodesCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")

	rootCmd.AddCommand(podsCmd)
	podsCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")

	rootCmd.AddCommand(secretsCmd)

	rootCmd.AddCommand(servicesCmd)
	servicesCmd.Flags().StringVarP(&output, "output", "o", "", "Output format.")

}

func resourceSearch(args []string, resourceType string) {
	var keyword string

	if len(args) >= 1 && args[0] != "" {
		keyword = utils.TrimQuoteAndSpace(args[0])
	}

	switch resourceType {
	// apps/v1
	case "daemonsets":
		resources.Daemonsets(searchOptions, keyword, output == "wide")
	case "deployments":
		resources.Deployments(searchOptions, keyword, output == "wide")
	case "replicasets":
		resources.Replicasets(searchOptions, keyword, output == "wide")
	case "statefulsets":
		resources.Statefulsets(searchOptions, keyword, output == "wide")

	// autoscaling/v1
	case "hpas":
		resources.Hpas(searchOptions, keyword)

	// batch/v1
	case "jobs":
		resources.Jobs(searchOptions, keyword)

	// networking.k8s.io/v1
	case "ingresses":
		resources.Ingresses(searchOptions, keyword)

	// storage.k8s.io/v1
	case "csidrivers":
		resources.CsiDrivers(searchOptions, keyword)
	case "storageclasses":
		resources.StorageClasses(searchOptions, keyword)

	// v1
	case "configmaps":
		resources.ConfigMaps(searchOptions, keyword)
	case "nodes":
		resources.Nodes(searchOptions, keyword, output == "wide")
	case "pods":
		resources.Pods(searchOptions, keyword, output == "wide")
	case "secrets":
		resources.Secrets(searchOptions, keyword)
	case "services":
		resources.Services(searchOptions, keyword, output == "wide")

	// default
	default:
		break
	}
}
