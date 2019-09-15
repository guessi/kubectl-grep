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

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, constants.StatefulsetHeaderWide)
	} else {
		fmt.Fprintln(w, constants.StatefulsetHeader)
	}

	for _, s := range statefulsetList.Items {
		// return all statefulsets under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(s.Name, keyword)
			if !match {
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
				s.Status.Replicas,
				s.Status.ReadyReplicas,
				age,
				strings.Join(names, ","),
				strings.Join(images, ","),
			)
		} else {
			statefulsetInfo = fmt.Sprintf(constants.StatefulsetRowTemplate,
				s.Namespace,
				s.Name,
				s.Status.Replicas,
				s.Status.ReadyReplicas,
				age,
			)
		}
		fmt.Fprintln(w, statefulsetInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
