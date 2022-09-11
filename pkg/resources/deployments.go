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

// Deployments - a public function for searching deployments with keyword
func Deployments(opt *options.SearchOptions, keyword string, wide bool) {
	var deploymentInfo string

	deploymentList := utils.DeploymentList(opt)

	if len(deploymentList.Items) <= 0 {
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

	if wide {
		fmt.Fprintln(w, constants.DeploymentHeaderWide)
	} else {
		fmt.Fprintln(w, constants.DeploymentHeader)
	}
	for _, d := range deploymentList.Items {
		// return all if no keyword specific
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
