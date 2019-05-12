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

// Deployments - a public function for searching deployments with keyword
func Deployments(namespace string, allNamespaces bool, selector, fieldSelector, keyword string, wide bool) {
	var deploymentInfo string

	deploymentList := utils.DeploymentList(namespace, allNamespaces, selector, fieldSelector)

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, constants.DeploymentHeaderWide)
	} else {
		fmt.Fprintln(w, constants.DeploymentHeader)
	}
	for _, d := range deploymentList.Items {
		// return all deployments under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(d.Name, keyword)
			if !match {
				continue
			}
		}

		age := utils.GetAge(time.Since(d.CreationTimestamp.Time))
		containers := d.Spec.Template.Spec.Containers

		if wide {
			names := []string{}
			images := []string{}

			for _, n := range containers {
				names = append(names, n.Name)
				images = append(images, n.Image)
			}

			deploymentInfo = fmt.Sprintf(constants.DeploymentRowTemplateWide,
				d.Namespace,
				d.Name,
				d.Status.Replicas,
				d.Status.ReadyReplicas,
				d.Status.UpdatedReplicas,
				d.Status.AvailableReplicas,
				age,
				strings.Join(names, ","),
				strings.Join(images, ","),
			)
		} else {
			deploymentInfo = fmt.Sprintf(constants.DeploymentRowTemplate,
				d.Namespace,
				d.Name,
				d.Status.Replicas,
				d.Status.ReadyReplicas,
				d.Status.UpdatedReplicas,
				d.Status.AvailableReplicas,
				age,
			)
		}
		fmt.Fprintln(w, deploymentInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
