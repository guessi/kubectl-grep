package options

import "time"

// TODO: integrate with "k8s.io/cli-runtime/pkg/genericclioptions"

type SearchOptions struct {
	AllNamespaces  bool
	Namespace      string
	Selector       string
	FieldSelector  string
	InvertMatch    bool
	ExcludePattern string
	Timeout        time.Duration
}

// NewSearchOptions - genericclioptions wrapper for searchOptions
func NewSearchOptions() *SearchOptions {
	return &SearchOptions{}
}
