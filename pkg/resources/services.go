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

// Services - a public function for searching services with keyword
func Services(ctx context.Context, opt *options.SearchOptions, keyword string, wide bool) error {
	var serviceInfo string

	serviceList, err := utils.ServiceList(ctx, opt)
	if err != nil {
		return err
	}

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
		return nil
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

		if !utils.MatchesKeyword(s.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(s.Name, opt.ExcludePattern) {
			continue
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
				if i.Hostname != "" {
					externalIPs = append(externalIPs, i.Hostname)
				} else if i.IP != "" {
					externalIPs = append(externalIPs, i.IP)
				}
			}
		} else {
			externalIPs = s.Spec.ExternalIPs
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

	return nil
}
