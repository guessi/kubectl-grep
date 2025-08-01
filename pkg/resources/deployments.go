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
)

// Deployments - a public function for searching deployments with keyword
func Deployments(ctx context.Context, opt *options.SearchOptions, keyword string, wide bool) error {
	var deploymentInfo string

	deploymentList, err := utils.DeploymentList(ctx, opt)
	if err != nil {
		return err
	}

	if len(deploymentList.Items) == 0 {
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
		fmt.Fprintln(w, constants.DeploymentHeaderWide)
	} else {
		fmt.Fprintln(w, constants.DeploymentHeader)
	}
	for _, d := range deploymentList.Items {
		if !utils.MatchesKeyword(d.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(d.Name, opt.ExcludePattern) {
			continue
		}

		age := utils.GetAge(time.Since(d.CreationTimestamp.Time))
		containers := d.Spec.Template.Spec.Containers

		var replicas int32 = 0
		if d.Spec.Replicas != nil {
			replicas = *d.Spec.Replicas
		}

		if wide {
			var names []string
			var images []string
			var selectors []string

			for _, n := range containers {
				names = append(names, n.Name)
				images = append(images, n.Image)
			}

			for k, v := range d.Spec.Selector.MatchLabels {
				selectors = append(selectors, fmt.Sprintf("%s=%s", k, v))
			}

			deploymentInfo = fmt.Sprintf(constants.DeploymentRowTemplateWide,
				d.Namespace,
				d.Name,
				d.Status.ReadyReplicas,
				replicas,
				d.Status.UpdatedReplicas,
				d.Status.AvailableReplicas,
				age,
				strings.Join(names, ","),
				strings.Join(images, ","),
				strings.Join(selectors, ","),
			)
		} else {
			deploymentInfo = fmt.Sprintf(constants.DeploymentRowTemplate,
				d.Namespace,
				d.Name,
				d.Status.ReadyReplicas,
				replicas,
				d.Status.UpdatedReplicas,
				d.Status.AvailableReplicas,
				age,
			)
		}
		fmt.Fprintln(w, deploymentInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())

	return nil
}
