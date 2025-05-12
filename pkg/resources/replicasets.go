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

	if len(replicasetList.Items) == 0 {
		ns := opt.Namespace
		if opt.AllNamespaces {
			fmt.Println("No resources found.")
		} else {
			if ns == "" {
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
		if !utils.MatchesKeyword(s.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(s.Name, opt.ExcludePattern) {
			continue
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))

		var replicas int32 = 0
		if s.Spec.Replicas != nil {
			replicas = *s.Spec.Replicas
		}

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
				replicas,
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
				replicas,
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
