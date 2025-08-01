package resources

import (
	"bytes"
	"context"
	"fmt"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-grep/pkg/constants"
	"github.com/guessi/kubectl-grep/pkg/options"
	"github.com/guessi/kubectl-grep/pkg/utils"
)

// Hpas - a public function for searching hpas with keyword
func Hpas(ctx context.Context, opt *options.SearchOptions, keyword string) error {
	hpaList, err := utils.HpaList(ctx, opt)
	if err != nil {
		return err
	}

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
		return nil
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.HpaHeader)
	for _, h := range hpaList.Items {
		if !utils.MatchesKeyword(h.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(h.Name, opt.ExcludePattern) {
			continue
		}

		var age string = utils.GetAge(time.Since(h.CreationTimestamp.Time))
		var currentCPUUtilizationPercentage string = "<unknown>"
		var targetCPUUtilizationPercentage string = "<unknown>"

		if h.Status.CurrentCPUUtilizationPercentage != nil {
			currentCPUUtilizationPercentage = fmt.Sprintf("%d%%", *h.Status.CurrentCPUUtilizationPercentage)
		}

		if h.Spec.TargetCPUUtilizationPercentage != nil {
			targetCPUUtilizationPercentage = fmt.Sprintf("%d%%", *h.Spec.TargetCPUUtilizationPercentage)
		}

		hpaInfo := fmt.Sprintf(constants.HpaRowTemplate,
			h.Namespace,
			h.Name,
			h.Spec.ScaleTargetRef.Kind,
			h.Spec.ScaleTargetRef.Name,
			currentCPUUtilizationPercentage,
			targetCPUUtilizationPercentage,
			*h.Spec.MinReplicas,
			h.Spec.MaxReplicas,
			h.Status.CurrentReplicas,
			age,
		)
		fmt.Fprintln(w, hpaInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())

	return nil
}
