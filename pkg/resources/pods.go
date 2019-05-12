package resources

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-search/pkg/constants"
	"github.com/guessi/kubectl-search/pkg/utils"
)

// Pods - a public function for searching pods with keyword
func Pods(namespace string, allNamespaces bool, selector, fieldSelector, keyword string, wide bool) {
	var podInfo string

	podList := utils.PodList(namespace, allNamespaces, selector, fieldSelector)

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, constants.PodHeaderWide)
	} else {
		fmt.Fprintln(w, constants.PodHeader)
	}
	for _, p := range podList.Items {
		// return all pods under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(p.Name, keyword)
			if !match {
				continue
			}
		}

		var containerCount int = len(p.Status.ContainerStatuses)
		var readyCount int32
		var restartCount int32

		for _, cs := range p.Status.ContainerStatuses {
			restartCount += cs.RestartCount
			if cs.Ready {
				readyCount++
			}
		}

		age, ageUnit := utils.GetAge(time.Since(p.CreationTimestamp.Time).Seconds())

		if wide {
			podInfo = fmt.Sprintf(constants.PodRowTemplateWide,
				p.Namespace,
				p.Name,
				readyCount, containerCount,
				p.Status.Phase,
				restartCount,
				age, ageUnit,
				p.Status.PodIP,
				p.Spec.NodeName,
			)
		} else {
			podInfo = fmt.Sprintf(constants.PodRowTemplate,
				p.Namespace,
				p.Name,
				readyCount, containerCount,
				p.Status.Phase,
				restartCount,
				age, ageUnit,
			)
		}
		fmt.Fprintln(w, podInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
