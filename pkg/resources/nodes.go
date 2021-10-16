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

// Nodes - a public function for searching nodes with keyword
func Nodes(opt *options.SearchOptions, keyword string, wide bool) {
	var nodeInfo string

	nodeList := utils.NodeList(opt)

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

		for label, _ := range n.Labels {
			if strings.HasPrefix(label, "node-role.kubernetes.io") {
				roles = append(roles, strings.SplitN(label, "/", 2)[1])
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

		age := utils.GetAge(time.Since(n.CreationTimestamp.Time))

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
				age,
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
				age,
				n.Status.NodeInfo.KubeletVersion,
			)
		}
		fmt.Fprintln(w, nodeInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
