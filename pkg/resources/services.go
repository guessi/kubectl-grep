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

// Services - a public function for searching services with keyword
func Services(opt *options.SearchOptions, keyword string, wide bool) {
	var serviceInfo string

	serviceList := utils.ServiceList(opt)

	if len(serviceList.Items) == 0 {
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
		fmt.Fprintln(w, constants.ServicesHeaderWide)
	} else {
		fmt.Fprintln(w, constants.ServicesHeader)
	}
	for _, s := range serviceList.Items {
		var ports []string
		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(s.Name, keyword)
			if match == opt.InvertMatch {
				continue
			}
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))

		for _, p := range s.Spec.Ports {
			var concatenated string
			if p.NodePort != 0 {
				concatenated = fmt.Sprintf("%d:%d/%s", p.Port, p.NodePort, p.Protocol)
			} else {
				concatenated = fmt.Sprintf("%d/%s", p.Port, p.Protocol)
			}
			ports = append(ports, concatenated)
		}

		var selectors []string
		var selector string
		if s.Spec.Selector != nil {
			for k, v := range s.Spec.Selector {
				selector = fmt.Sprintf("%s=%s", k, v)
				selectors = append(selectors, selector)
			}
		}
		selectorOutput := "<none>"
		if len(selectors) > 0 {
			selectorOutput = strings.Join(selectors, ",")
		}

		var externalIPs []string
		if s.Spec.ExternalIPs == nil {
			for _, i := range s.Status.LoadBalancer.Ingress {
				externalIPs = append(externalIPs, i.Hostname)
			}
		}

		var externalIPsDisplay string = "<none>"
		if len(externalIPs) > 0 {
			externalIPsDisplay = strings.Join(externalIPs, ",")
		}

		if wide {
			serviceInfo = fmt.Sprintf(constants.ServicesRowTemplateWide,
				s.Namespace,
				s.Name,
				s.Spec.Type,
				s.Spec.ClusterIP,
				externalIPsDisplay,
				strings.Join(ports, ","),
				age,
				selectorOutput,
			)
		} else {
			serviceInfo = fmt.Sprintf(constants.ServicesRowTemplate,
				s.Namespace,
				s.Name,
				s.Spec.Type,
				s.Spec.ClusterIP,
				externalIPsDisplay,
				strings.Join(ports, ","),
				age,
			)
		}

		fmt.Fprintln(w, serviceInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
