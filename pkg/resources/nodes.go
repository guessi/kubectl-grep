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
	v1 "k8s.io/api/core/v1"
)

// Nodes - a public function for searching nodes with keyword
func Nodes(ctx context.Context, opt *options.SearchOptions, keyword string, wide bool) error {
	var nodeInfo string

	nodeList, err := utils.NodeList(ctx, opt)
	if err != nil {
		return err
	}

	if len(nodeList.Items) == 0 {
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
		fmt.Fprintln(w, constants.NodeHeaderWide)
	} else {
		fmt.Fprintln(w, constants.NodeHeader)
	}
	for _, n := range nodeList.Items {
		if !utils.MatchesKeyword(n.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(n.Name, opt.ExcludePattern) {
			continue
		}

		var roles []string

		for label := range n.Labels {
			if strings.HasPrefix(label, "node-role.kubernetes.io") {
				roles = append(roles, strings.SplitN(label, "/", 2)[1])
			}
		}
		if len(roles) <= 0 {
			roles = append(roles, "<none>")
		}

		var nodeStatus string = "Unknown"
		for _, condition := range n.Status.Conditions {
			if condition.Type == v1.NodeReady {
				if condition.Status == v1.ConditionTrue {
					nodeStatus = "Ready"
				} else {
					nodeStatus = "NotReady"
				}
			}
		}

		if n.Spec.Unschedulable {
			nodeStatus = nodeStatus + ",SchedulingDisabled"
		}

		age := utils.GetAge(time.Since(n.CreationTimestamp.Time))

		if wide {
			var extAddr string = "<none>"
			var intAddr string = "<none>"

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

	return nil
}
