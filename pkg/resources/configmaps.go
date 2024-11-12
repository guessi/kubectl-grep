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

// ConfigMaps - a public function for searching configmaps with keyword
func ConfigMaps(opt *options.SearchOptions, keyword string) {
	var configMapInfo string

	configMapList := utils.ConfigMapList(opt)

	if len(configMapList.Items) == 0 {
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

	fmt.Fprintln(w, constants.ConfigMapHeader)

	for _, cm := range configMapList.Items {
		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(cm.Name, keyword)
			if match == opt.InvertMatch {
				continue
			}
		}

		age := utils.GetAge(time.Since(cm.CreationTimestamp.Time))

		configMapInfo = fmt.Sprintf(constants.ConfigMapRowTemplate,
			cm.Namespace,
			cm.Name,
			len(cm.Data),
			age,
		)

		fmt.Fprintln(w, configMapInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
