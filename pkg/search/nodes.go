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

// Nodes - a public function for searching nodes with keyword
func Nodes(selector, fieldSelector, keyword string, wide bool) {
	var nodeInfo string

	nodeList := utils.NodeList(selector, fieldSelector)

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, constants.NodeHeaderWide)
	} else {
		fmt.Fprintln(w, constants.NodeHeader)
	}
	for _, n := range nodeList.Items {
		// return all pods under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(n.Name, keyword)
			if !match {
				continue
			}
		}

		var roles []string
		var nodeStatus string

		for _, label := range n.Labels {
			// FIXME: should detect other type of roles
			if strings.Contains(label, "node-role.kubernetes.io/master") {
				roles = append(roles, "master")
			}
		}
		if len(roles) <= 0 {
			roles = append(roles, "<none>")
		}

		if !n.Spec.Unschedulable {
			nodeStatus = "Ready"
		} else {
			nodeStatus = "Ready,SchedulingDisabled"
		}

		age, ageUnit := utils.GetAge(time.Since(n.CreationTimestamp.Time).Seconds())

		if wide {
			var extAddr string
			var intAddr string

			for _, addr := range n.Status.Addresses {
				if addr.Type == "ExternalIP" {
					extAddr = addr.Address
				}
				if addr.Type == "InternalIP" {
					intAddr = addr.Address
				}
			}

			nodeInfo = fmt.Sprintf(constants.NodeRowTemplateWide,
				n.Name,
				nodeStatus,
				strings.Join(roles, ","),
				age, ageUnit,
				n.Status.NodeInfo.KubeletVersion,
				intAddr,
				extAddr,
				n.Status.NodeInfo.OSImage,
				n.Status.NodeInfo.KernelVersion,
				n.Status.NodeInfo.ContainerRuntimeVersion,
			)
		} else {
			nodeInfo = fmt.Sprintf(constants.NodeRowTemplate,
				n.Name,
				nodeStatus,
				strings.Join(roles, ","),
				age, ageUnit,
				n.Status.NodeInfo.KubeletVersion,
			)
		}
		fmt.Fprintln(w, nodeInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
