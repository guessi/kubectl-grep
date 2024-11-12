package resources

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-grep/pkg/constants"
	"github.com/guessi/kubectl-grep/pkg/options"
	"github.com/guessi/kubectl-grep/pkg/utils"
)

// Replicasets - a public function for searching Replicasets with keyword
func Replicasets(opt *options.SearchOptions, keyword string, wide bool) {
	var replicasetInfo string

	replicasetList := utils.ReplicaSetList(opt)

	if len(replicasetList.Items) <= 0 {
		if opt.AllNamespaces {
			fmt.Printf("No resources found.\n")
		} else {
			var ns = opt.Namespace
			if len(opt.Namespace) <= 0 {
				ns = "default"
			}
			fmt.Printf("No resources found in %s namespace.\n", ns)
		}
		return
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, constants.ReplicasetHeaderWide)
	} else {
		fmt.Fprintln(w, constants.ReplicasetHeader)
	}

	for _, s := range replicasetList.Items {
		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(s.Name, keyword)
			if match == opt.InvertMatch {
				continue
			}
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))

		if wide {
			names := []string{}
			images := []string{}

			for _, n := range s.Spec.Template.Spec.Containers {
				names = append(names, n.Name)
				images = append(images, n.Image)
			}

			selectors := []string{}
			for k, v := range s.Spec.Selector.MatchLabels {
				selectors = append(selectors, fmt.Sprintf("%s=%s", k, v))
			}

			replicasetInfo = fmt.Sprintf(constants.ReplicasetRowTemplateWide,
				s.Namespace,
				s.Name,
				*s.Spec.Replicas,
				s.Status.Replicas,
				s.Status.ReadyReplicas,
				age,
				strings.Join(names, ","),
				strings.Join(images, ","),
				strings.Join(selectors, ","),
			)
		} else {
			replicasetInfo = fmt.Sprintf(constants.ReplicasetRowTemplate,
				s.Namespace,
				s.Name,
				*s.Spec.Replicas,
				s.Status.Replicas,
				s.Status.ReadyReplicas,
				age,
			)
		}
		fmt.Fprintln(w, replicasetInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
