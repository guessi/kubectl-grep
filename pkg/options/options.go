package options

// TODO: integrate with "k8s.io/cli-runtime/pkg/genericclioptions"

type SearchOptions struct {
	AllNamespaces bool
	Namespace     string
	Selector      string
	FieldSelector string
	InvertMatch   bool
}

// NewSearchOptions - genericclioptions wrapper for searchOptions
func NewSearchOptions() *SearchOptions {
	return &SearchOptions{}
}
