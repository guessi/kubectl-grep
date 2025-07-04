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

// Pods - a public function for searching pods with keyword
func Pods(ctx context.Context, opt *options.SearchOptions, keyword string, wide bool) error {
	var podInfo string

	podList, err := utils.PodList(ctx, opt)
	if err != nil {
		return err
	}

	if len(podList.Items) == 0 {
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

	if wide {
		fmt.Fprintln(w, constants.PodHeaderWide)
	} else {
		fmt.Fprintln(w, constants.PodHeader)
	}
	for _, p := range podList.Items {
		if !utils.MatchesKeyword(p.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(p.Name, opt.ExcludePattern) {
			continue
		}

		var containerCount int = len(p.Spec.Containers)
		var readyCount int32
		var restartCount int32

		for _, cs := range p.Status.ContainerStatuses {
			restartCount += cs.RestartCount
			if cs.Ready {
				readyCount++
			}
		}

		var podIP string = "<none>"
		if len(p.Status.PodIP) > 0 {
			podIP = p.Status.PodIP
		}

		var nodeName string = "<none>"
		if len(p.Spec.NodeName) > 0 {
			nodeName = p.Spec.NodeName
		}

		age := utils.GetAge(time.Since(p.CreationTimestamp.Time))

		if wide {
			podInfo = fmt.Sprintf(constants.PodRowTemplateWide,
				p.Namespace,
				p.Name,
				readyCount, containerCount,
				p.Status.Phase,
				restartCount,
				age,
				podIP,
				nodeName,
			)
		} else {
			podInfo = fmt.Sprintf(constants.PodRowTemplate,
				p.Namespace,
				p.Name,
				readyCount, containerCount,
				p.Status.Phase,
				restartCount,
				age,
			)
		}
		fmt.Fprintln(w, podInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())

	return nil
}
