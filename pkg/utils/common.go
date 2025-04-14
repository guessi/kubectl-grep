package utils

import (
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/duration"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// setOptions - set common options for clientset
func setOptions(opt *options.SearchOptions) (string, *metav1.ListOptions) {
	// set default namespace as "default"
	namespace := "default"

	// override `namespace` if `--all-namespaces` exist
	if opt.AllNamespaces {
		namespace = ""
	} else {
		if len(opt.Namespace) > 0 {
			namespace = opt.Namespace
		} else {
			ns, _, err := client.ClientConfig().Namespace()
			if err != nil {
				log.WithFields(log.Fields{
					"err": err.Error(),
				}).Debug("Failed to resolve namespace")
			} else {
				namespace = ns
			}
		}
	}

	// retrieve listOptions from meta
	listOptions := &metav1.ListOptions{
		LabelSelector: opt.Selector,
		FieldSelector: opt.FieldSelector,
	}
	return namespace, listOptions
}

// TrimQuoteAndSpace - remove Spaces, Tabs, SingleQuotes, DoubleQuites
func TrimQuoteAndSpace(input string) string {
	if len(input) >= 2 {
		if input[0] == '"' && input[len(input)-1] == '"' {
			return input[1 : len(input)-1]
		}
		if input[0] == '\'' && input[len(input)-1] == '\'' {
			return input[1 : len(input)-1]
		}
	}
	return strings.TrimSpace(input)
}

// GetAge - return human readable time expression
func GetAge(d time.Duration) string {
	return duration.HumanDuration(d)
}

// BoolValue
func BoolValue(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

// BoolString
func BoolString(b *bool) string {
	return strconv.FormatBool(BoolValue(b))
}

func MatchesKeyword(target string, keyword string, invertMatch bool) bool {
	if len(keyword) == 0 {
		return true
	}
	match := strings.Contains(target, keyword)
	return match != invertMatch
}

func ShouldExcludeResource(target string, excludePattern string) bool {
	excludePattern = strings.TrimSpace(excludePattern)
	if excludePattern == "" {
		return false
	}

	for _, exclude := range strings.Split(excludePattern, ",") {
		exclude = strings.TrimSpace(exclude)
		if exclude != "" && strings.Contains(target, exclude) {
			return true
		}
	}
	return false
}
