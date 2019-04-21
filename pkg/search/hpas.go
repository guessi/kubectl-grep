package search

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-search/pkg/search/constants"
	"github.com/guessi/kubectl-search/pkg/search/utils"
)

// Hpas - a public function for searching hpas with keyword
func Hpas(namespace string, allNamespaces bool, selector, filedSelector, keyword string) {
	hpaList := utils.HpaList(namespace, allNamespaces, selector, filedSelector)

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

		age, ageUnit := utils.GetAge(time.Since(h.CreationTimestamp.Time).Seconds())

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
			age, ageUnit,
		)
		fmt.Fprintln(w, hpaInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
