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

// Secrets - a public function for searching secrets with keyword
func Secrets(ctx context.Context, opt *options.SearchOptions, keyword string) error {
	var secretInfo string

	secretList, err := utils.SecretList(ctx, opt)
	if err != nil {
		return err
	}

	if len(secretList.Items) == 0 {
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

	fmt.Fprintln(w, constants.SecretHeader)

	for _, s := range secretList.Items {
		if !utils.MatchesKeyword(s.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(s.Name, opt.ExcludePattern) {
			continue
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))

		secretInfo = fmt.Sprintf(constants.SecretRowTemplate,
			s.Namespace,
			s.Name,
			s.Type,
			len(s.Data),
			age,
		)

		fmt.Fprintln(w, secretInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())

	return nil
}
