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

// Hpas - a public function for searching hpas with keyword
func Hpas(opt *options.SearchOptions, keyword string) {
	hpaList := utils.HpaList(opt)

	if len(hpaList.Items) == 0 {
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

	fmt.Fprintln(w, constants.HpaHeader)
	for _, h := range hpaList.Items {
		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(h.Name, keyword)
			if match == opt.InvertMatch {
				continue
			}
		}

		var age string = utils.GetAge(time.Since(h.CreationTimestamp.Time))
		var currentCPUUtilizationPercentage string = "<unknown>"

		if h.Status.CurrentCPUUtilizationPercentage != nil {
			currentCPUUtilizationPercentage = fmt.Sprintf("%d%%", *h.Status.CurrentCPUUtilizationPercentage)
		}

		hpaInfo := fmt.Sprintf(constants.HpaRowTemplate,
			h.Namespace,
			h.Name,
			h.Spec.ScaleTargetRef.Kind,
			h.Spec.ScaleTargetRef.Name,
			currentCPUUtilizationPercentage,
			*h.Spec.TargetCPUUtilizationPercentage,
			*h.Spec.MinReplicas,
			h.Spec.MaxReplicas,
			h.Status.CurrentReplicas,
			age,
		)
		fmt.Fprintln(w, hpaInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
