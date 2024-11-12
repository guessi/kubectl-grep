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

// ClusterRoleBindings - a public function for searching clusterrolebindings with keyword
func ClusterRoleBindings(opt *options.SearchOptions, keyword string) {
	var clusterRoleBindingInfo string

	clusterRoleBindingList := utils.ClusterRoleBindingList(opt)

	if len(clusterRoleBindingList.Items) <= 0 {
		fmt.Printf("No resources found.\n")
		return
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.ClusterRoleBindingsHeader)

	for _, c := range clusterRoleBindingList.Items {
		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(c.Name, keyword)
			if match == opt.InvertMatch {
				continue
			}
		}

		age := utils.GetAge(time.Since(c.CreationTimestamp.Time))

		clusterRoleBindingInfo = fmt.Sprintf(constants.ClusterRoleBindingsRowTemplate,
			c.Name,
			"ClusterRole/"+c.RoleRef.Name,
			age,
		)

		fmt.Fprintln(w, clusterRoleBindingInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
