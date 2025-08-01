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

// ClusterRoleBindings - a public function for searching clusterrolebindings with keyword
func ClusterRoleBindings(ctx context.Context, opt *options.SearchOptions, keyword string) error {
	var clusterRoleBindingInfo string

	clusterRoleBindingList, err := utils.ClusterRoleBindingList(ctx, opt)
	if err != nil {
		return err
	}

	if len(clusterRoleBindingList.Items) <= 0 {
		fmt.Printf("No resources found.\n")
		return nil
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.ClusterRoleBindingsHeader)

	for _, c := range clusterRoleBindingList.Items {
		if !utils.MatchesKeyword(c.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(c.Name, opt.ExcludePattern) {
			continue
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

	return nil
}
