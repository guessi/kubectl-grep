package resources

import (
	"bytes"
	"context"
	"fmt"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-grep/pkg/constants"
	"github.com/guessi/kubectl-grep/pkg/options"
	"github.com/guessi/kubectl-grep/pkg/utils"
)

// ClusterRoles - a public function for searching clusterroles with keyword
func ClusterRoles(ctx context.Context, opt *options.SearchOptions, keyword string) error {
	var clusterRoleInfo string

	clusterRoleList, err := utils.ClusterRoleList(ctx, opt)
	if err != nil {
		return err
	}

	if len(clusterRoleList.Items) <= 0 {
		fmt.Printf("No resources found.\n")
		return nil
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

	return nil
}
