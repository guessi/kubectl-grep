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
	autoscalingv2 "k8s.io/api/autoscaling/v2"
)

// Hpas - a public function for searching hpas with keyword
func Hpas(opt *options.SearchOptions, keyword string) {
	hpaList := utils.HpaList(opt)

	if len(hpaList.Items) <= 0 {
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

	fmt.Fprintln(w, constants.HpaHeader)
	for _, h := range hpaList.Items {
		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(h.Name, keyword)
			if !match {
				continue
			}
		}

		var age string = utils.GetAge(time.Since(h.CreationTimestamp.Time))

		var targetCPUUtilizationPercentage string = "<unknown>"
		var currentCPUUtilizationPercentage string = "<unknown>"

		// var targetMemoryUtilizationPercentage string = "<unknown>"
		// var currentMemoryUtilizationPercentage string = "<unknown>"

		for _, metric := range h.Spec.Metrics {
			if metric.Type == autoscalingv2.ResourceMetricSourceType {
				if metric.Resource.Name == "cpu" {
					targetCPUUtilizationPercentage = fmt.Sprintf("%d%%", *metric.Resource.Target.AverageUtilization)
				}
				// if metric.Resource.Name == "memory" {
				// 	targetMemoryUtilizationPercentage = fmt.Sprintf("%d%%", *metric.Resource.Target.AverageUtilization)
				// }
			}
		}

		for _, metric := range h.Status.CurrentMetrics {
			if metric.Type == autoscalingv2.ResourceMetricSourceType {
				if metric.Resource.Name == "cpu" {
					currentCPUUtilizationPercentage = fmt.Sprintf("%d%%", *metric.Resource.Current.AverageUtilization)
				}
				// if metric.Resource.Name == "memory" {
				// 	currentMemoryUtilizationPercentage = fmt.Sprintf("%d%%", *metric.Resource.Current.AverageUtilization)
				// }
			}
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
}
