package resources

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-grep/pkg/constants"
	"github.com/guessi/kubectl-grep/pkg/options"
	"github.com/guessi/kubectl-grep/pkg/utils"
)

// RoleBindings - a public function for searching rolebindings with keyword
func RoleBindings(opt *options.SearchOptions, keyword string) {
	var roleBindingInfo string

	roleBindingList := utils.RoleBindingList(opt)

	if len(roleBindingList.Items) == 0 {
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

	fmt.Fprintln(w, constants.RoleBindingsHeader)

	for _, r := range roleBindingList.Items {
		if !utils.MatchesKeyword(r.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(r.Name, opt.ExcludePattern) {
			continue
		}

		age := utils.GetAge(time.Since(r.CreationTimestamp.Time))

		roleBindingInfo = fmt.Sprintf(constants.RoleBindingsRowTemplate,
			r.Namespace,
			r.Name,
			"Role/"+r.RoleRef.Name,
			age,
		)

		fmt.Fprintln(w, roleBindingInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
