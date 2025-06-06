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

// ServiceAccounts - a public function for searching serviceaccounts with keyword
func ServiceAccounts(opt *options.SearchOptions, keyword string) {
	var serviceAccountInfo string

	serviceAccountList := utils.ServiceAccountList(opt)

	if len(serviceAccountList.Items) == 0 {
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

	fmt.Fprintln(w, constants.ServiceAccountsHeader)

	for _, s := range serviceAccountList.Items {
		if !utils.MatchesKeyword(s.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(s.Name, opt.ExcludePattern) {
			continue
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))

		serviceAccountInfo = fmt.Sprintf(constants.ServiceAccountsRowTemplate,
			s.Namespace,
			s.Name,
			len(s.Secrets),
			age,
		)

		fmt.Fprintln(w, serviceAccountInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
