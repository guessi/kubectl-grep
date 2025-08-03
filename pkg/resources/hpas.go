package resources

import (
	"bytes"
	"context"
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

		var targetCPUUtilization string = constants.UNKNOWN
		var targetMemoryUtilization string = constants.UNKNOWN

		var currentCPUUtilization string = constants.UNKNOWN
		var currentMemoryUtilization string = constants.UNKNOWN

		// Process target metrics
		for _, metric := range h.Spec.Metrics {
			if metric.Type == autoscalingv2.ResourceMetricSourceType {
				switch metric.Resource.Name {
				case "cpu":
					targetCPUUtilization = utils.FormatUtilization(metric.Resource.Target.AverageUtilization)
				case "memory":
					targetMemoryUtilization = utils.FormatUtilization(metric.Resource.Target.AverageUtilization)
				}
			}
		}

		// Process current metrics
		for _, metric := range h.Status.CurrentMetrics {
			if metric.Type == autoscalingv2.ResourceMetricSourceType {
				switch metric.Resource.Name {
				case "cpu":
					currentCPUUtilization = utils.FormatUtilization(metric.Resource.Current.AverageUtilization)
				case "memory":
					currentMemoryUtilization = utils.FormatUtilization(metric.Resource.Current.AverageUtilization)
				}
			}
		}

		var targetsFieldInfo string
		var metrics []string

		if targetCPUUtilization != constants.UNKNOWN {
			metrics = append(metrics, fmt.Sprintf("cpu: %s/%s", currentCPUUtilization, targetCPUUtilization))
		}
		if targetMemoryUtilization != constants.UNKNOWN {
			metrics = append(metrics, fmt.Sprintf("memory: %s/%s", currentMemoryUtilization, targetMemoryUtilization))
		}
		if len(metrics) > 0 {
			targetsFieldInfo = fmt.Sprintf("%s", strings.Join(metrics, ", "))
		}

		hpaInfo := fmt.Sprintf(constants.HpaRowTemplate,
			h.Namespace,
			h.Name,
			h.Spec.ScaleTargetRef.Kind,
			h.Spec.ScaleTargetRef.Name,
			targetsFieldInfo,
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
