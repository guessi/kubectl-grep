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

// Roles - a public function for searching roles with keyword
func Roles(opt *options.SearchOptions, keyword string) {
	var roleInfo string

	roleList := utils.RoleList(opt)

	if len(roleList.Items) == 0 {
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

	fmt.Fprintln(w, constants.RolesHeader)

	for _, r := range roleList.Items {
		if !utils.MatchesKeyword(r.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(r.Name, opt.ExcludePattern) {
			continue
		}

		createdAt := r.CreationTimestamp.Time

		roleInfo = fmt.Sprintf(constants.RolesRowTemplate,
			r.Namespace,
			r.Name,
			createdAt.UTC().Format(time.RFC3339),
		)

		fmt.Fprintln(w, roleInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
