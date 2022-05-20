package resources

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-grep/pkg/constants"
	"github.com/guessi/kubectl-grep/pkg/options"
	"github.com/guessi/kubectl-grep/pkg/utils"
)

// Ingresses - a public function for searching ingresses with keyword
func Ingresses(opt *options.SearchOptions, keyword string) {
	var ingressInfo string

	ingressList := utils.IngressList(opt)

	if len(ingressList.Items) <= 0 {
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

	fmt.Fprintln(w, constants.IngressHeader)

	for _, i := range ingressList.Items {
		var ingressClassName string
		var hosts, ports, addresses []string

		// return all ingresses under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(i.Name, keyword)
			if !match {
				continue
			}
		}

		age := utils.GetAge(time.Since(i.CreationTimestamp.Time))

		if i.Spec.IngressClassName == nil {
			ingressClassName = "<none>"
		} else {
			ingressClassName = *i.Spec.IngressClassName
		}

		for _, irs := range i.Spec.Rules {
			if len(irs.Host) > 0 {
				hosts = append(hosts, irs.Host)
			}

			for _, ips := range irs.IngressRuleValue.HTTP.Paths {
				ports = append(ports, strconv.Itoa(int(ips.Backend.Service.Port.Number)))
			}
		}

		for _, lbi := range i.Status.LoadBalancer.Ingress {
			addresses = append(addresses, lbi.IP)
		}

		ingressInfo = fmt.Sprintf(constants.IngressRowTemplate,
			i.Namespace,
			i.Name,
			ingressClassName,
			strings.Join(hosts, ","),
			strings.Join(addresses, ","),
			strings.Join(ports, ","),
			age,
		)

		fmt.Fprintln(w, ingressInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
