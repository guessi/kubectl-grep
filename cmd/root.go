package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/guessi/kubectl-grep/pkg/options"
)

var (
	output                  string
	rootCmdDescriptionShort = "Filter Kubernetes resources by matching their names"
	rootCmdDescriptionLong  = `Filter Kubernetes resources by matching their names

More info: https://github.com/guessi/kubectl-grep
`

	rootCmdExamples = `
List all pods in default namespace
$ kubectl grep pods

List all pods in all namespaces
$ kubectl grep pods -A

List all pods in namespace "start-lab" which contains keyword "flash"
$ kubectl grep pods -n star-lab flash
`
)

// generic search options handler
var searchOptions = options.NewSearchOptions()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "kubectl-grep",
	Short:   rootCmdDescriptionShort,
	Long:    rootCmdDescriptionLong,
	Example: rootCmdExamples,
}

func init() {
	// Global Flags
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.Namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	rootCmd.PersistentFlags().BoolVarP(
		&searchOptions.AllNamespaces, "all-namespaces", "A", false,
		"If present, list the requested object(s) across all namespaces.")
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.Selector, "selector", "l", "",
		"Selector (label query) to filter on. (e.g. -l key1=value1,key2=value2)")
	rootCmd.PersistentFlags().StringVar(
		&searchOptions.FieldSelector, "field-selector", "",
		"Selector (field query) to filter on. (e.g. --field-selector key1=value1,key2=value2)")
	rootCmd.PersistentFlags().BoolVarP(
		&searchOptions.InvertMatch, "invert-match", "v", false,
		"If present, filter out those not matching the specified patterns")
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.ExcludePattern, "exclude", "x", "",
		"If present, exclude those with specified pattern (comma-separated string)")
	rootCmd.PersistentFlags().DurationVarP(
		&searchOptions.Timeout, "timeout", "t", 30*time.Second,
		"Timeout for Kubernetes API calls (default: 30s)")
}

// createContextWithTimeout creates a context with timeout and cancellation support
func createContextWithTimeout() (context.Context, context.CancelFunc) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), searchOptions.Timeout)

	// Handle interrupt signals for graceful cancellation
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case <-c:
			fmt.Fprintln(os.Stderr, "\nOperation cancelled by user")
			cancel()
		case <-ctx.Done():
			// Context already done, cleanup signal handler
		}
		signal.Stop(c)
	}()

	return ctx, cancel
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
