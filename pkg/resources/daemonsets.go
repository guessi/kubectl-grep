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

// Daemonsets - a public function for searching daemonsets with keyword
func Daemonsets(opt *options.SearchOptions, keyword string, wide bool) {
	var daemonsetInfo string

	daemonsetList := utils.DaemonsetList(opt)

	if len(daemonsetList.Items) == 0 {
		ns := opt.Namespace
		if opt.AllNamespaces {
			fmt.Println("No resources found.")
		} else {
			if ns == "" {
				ns = "default"
			}
			fmt.Printf("No resources found in %s namespace.\n", ns)
		}
		return
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, constants.DaemonsetHeaderWide)
	} else {
		fmt.Fprintln(w, constants.DaemonsetHeader)
	}

	for _, d := range daemonsetList.Items {
		if !utils.MatchesKeyword(d.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(d.Name, opt.ExcludePattern) {
			continue
		}

		age := utils.GetAge(time.Since(d.CreationTimestamp.Time))
		containers := d.Spec.Template.Spec.Containers

		var nodeSelectors []string
		var nodeSelector string
		if d.Spec.Template.Spec.NodeSelector != nil {
			for k, v := range d.Spec.Template.Spec.NodeSelector {
				nodeSelector = fmt.Sprintf("%s=%s", k, v)
				nodeSelectors = append(nodeSelectors, nodeSelector)
			}
		}
		nodeSelectorOutput := "<none>"
		if len(nodeSelectors) > 0 {
			nodeSelectorOutput = strings.Join(nodeSelectors, ",")
		}

		var selectors []string
		var selector string
		if d.Spec.Selector.MatchLabels != nil {
			for k, v := range d.Spec.Selector.MatchLabels {
				selector = fmt.Sprintf("%s=%s", k, v)
				selectors = append(selectors, selector)
			}
		}
		selectorOutput := "<none>"
		if len(selectors) > 0 {
			selectorOutput = strings.Join(selectors, ",")
		}

		if wide {
			names := []string{}
			images := []string{}

			for _, n := range containers {
				names = append(names, n.Name)
				images = append(images, n.Image)
			}

			daemonsetInfo = fmt.Sprintf(constants.DaemonsetRowTemplateWide,
				d.Namespace,
				d.Name,
				d.Status.DesiredNumberScheduled,
				d.Status.CurrentNumberScheduled,
				d.Status.NumberReady,
				d.Status.UpdatedNumberScheduled,
				d.Status.NumberAvailable,
				nodeSelectorOutput,
				age,
				strings.Join(names, ","),
				strings.Join(images, ","),
				selectorOutput,
			)
		} else {
			daemonsetInfo = fmt.Sprintf(constants.DaemonsetRowTemplate,
				d.Namespace,
				d.Name,
				d.Status.DesiredNumberScheduled,
				d.Status.CurrentNumberScheduled,
				d.Status.NumberReady,
				d.Status.UpdatedNumberScheduled,
				d.Status.NumberAvailable,
				nodeSelectorOutput,
				age,
			)
		}
		fmt.Fprintln(w, daemonsetInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
