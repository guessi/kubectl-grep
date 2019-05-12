package resources

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-search/pkg/constants"
	"github.com/guessi/kubectl-search/pkg/options"
	"github.com/guessi/kubectl-search/pkg/utils"
)

// Hpas - a public function for searching hpas with keyword
func Hpas(opt *options.SearchOptions, keyword string) {
	hpaList := utils.HpaList(opt)

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.HpaHeader)
	for _, h := range hpaList.Items {
		// return all hpas under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(h.Name, keyword)
			if !match {
				continue
			}
		}

		age := utils.GetAge(time.Since(h.CreationTimestamp.Time))

		hpaInfo := fmt.Sprintf(constants.HpaRowTemplate,
			h.Namespace,
			h.Name,
			h.Spec.ScaleTargetRef.Kind,
			h.Spec.ScaleTargetRef.Name,
			*h.Status.CurrentCPUUtilizationPercentage,
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
