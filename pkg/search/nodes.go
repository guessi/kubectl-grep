package search

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	client "github.com/guessi/kubectl-search/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	nodesFields     = "NAME\tSTATUS\tROLES\tAGE\tVERSION"
	nodesFieldsWide = "NAME\tSTATUS\tROLES\tAGE\tVERSION\tINTERNAL-IP\tEXTERNAL-IP\tOS-IMAGE\tKERNEL-VERSION\tCONTAINER-RUNTIME"
	nInfo           string
)

// Pods - a public function for searching pods with keyword
func Nodes(selector, fieldSelector, keyword string, wide bool) {
	clientset := client.InitClient()

	listOptions := &metav1.ListOptions{
		LabelSelector: selector,
		FieldSelector: fieldSelector,
	}

	nodes, err := clientset.CoreV1().Nodes().List(*listOptions)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, nodesFieldsWide)
	} else {
		fmt.Fprintln(w, nodesFields)
	}
	for _, n := range nodes.Items {
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

		age, ageUnit := getAge(time.Since(n.CreationTimestamp.Time).Seconds())

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

			nInfo = fmt.Sprintf("%s\t%s\t%s\t%d%s\t%s\t%s\t%s\t%s\t%s\t%s",
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
			nInfo = fmt.Sprintf("%s\t%s\t%s\t%d%s\t%s",
				n.Name,
				nodeStatus,
				strings.Join(roles, ","),
				age, ageUnit,
				n.Status.NodeInfo.KubeletVersion,
			)
		}
		fmt.Fprintln(w, nInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
