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

// Statefulsets - a public function for searching Statefulsets with keyword
func Statefulsets(opt *options.SearchOptions, keyword string, wide bool) {
	var statefulsetInfo string

	statefulsetList := utils.StatefulSetList(opt)

	if len(statefulsetList.Items) <= 0 {
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
		fmt.Fprintln(w, constants.StatefulsetHeaderWide)
	} else {
		fmt.Fprintln(w, constants.StatefulsetHeader)
	}

	for _, s := range statefulsetList.Items {
		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(s.Name, keyword)
			if match == opt.InvertMatch {
				continue
			}
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))
		containers := s.Spec.Template.Spec.Containers

		if wide {
			names := []string{}
			images := []string{}

			for _, n := range containers {
				names = append(names, n.Name)
				images = append(images, n.Image)
			}

			statefulsetInfo = fmt.Sprintf(constants.StatefulsetRowTemplateWide,
				s.Namespace,
				s.Name,
				s.Status.ReadyReplicas,
				*s.Spec.Replicas,
				age,
				strings.Join(names, ","),
				strings.Join(images, ","),
			)
		} else {
			statefulsetInfo = fmt.Sprintf(constants.StatefulsetRowTemplate,
				s.Namespace,
				s.Name,
				s.Status.ReadyReplicas,
				*s.Spec.Replicas,
				age,
			)
		}
		fmt.Fprintln(w, statefulsetInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
