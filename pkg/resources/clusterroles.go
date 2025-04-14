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

// ClusterRoles - a public function for searching clusterroles with keyword
func ClusterRoles(opt *options.SearchOptions, keyword string) {
	var clusterRoleInfo string

	clusterRoleList := utils.ClusterRoleList(opt)

	if len(clusterRoleList.Items) <= 0 {
		fmt.Printf("No resources found.\n")
		return
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.ClusterRolesHeader)

	for _, c := range clusterRoleList.Items {
		if !utils.MatchesKeyword(c.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(c.Name, opt.ExcludePattern) {
			continue
		}

		createdAt := c.CreationTimestamp.Time

		clusterRoleInfo = fmt.Sprintf(constants.ClusterRolesRowTemplate,
			c.Name,
			createdAt.UTC().Format(time.RFC3339),
		)

		fmt.Fprintln(w, clusterRoleInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
