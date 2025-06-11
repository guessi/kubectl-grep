package resources

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/guessi/kubectl-grep/pkg/constants"
	"github.com/guessi/kubectl-grep/pkg/options"
	"github.com/guessi/kubectl-grep/pkg/utils"
)

// CsiDrivers - a public function for searching csidrivers with keyword
func CsiDrivers(ctx context.Context, opt *options.SearchOptions, keyword string) error {
	var csiDriverInfo string

	csiDriverList, err := utils.CsiDriverList(ctx, opt)
	if err != nil {
		return err
	}

	if len(csiDriverList.Items) == 0 {
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

	fmt.Fprintln(w, constants.CsiDriversHeader)

	for _, s := range csiDriverList.Items {
		if !utils.MatchesKeyword(s.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(s.Name, opt.ExcludePattern) {
			continue
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))

		var tokenRequest string
		if len(s.Spec.TokenRequests) <= 0 {
			tokenRequest = "<unset>"
		}

		var modes []string
		for _, m := range s.Spec.VolumeLifecycleModes {
			modes = append(modes, string(m))
		}

		csiDriverInfo = fmt.Sprintf(constants.CsiDriversRowTemplate,
			s.Name,
			utils.BoolString(s.Spec.AttachRequired),
			utils.BoolString(s.Spec.PodInfoOnMount),
			utils.BoolString(s.Spec.StorageCapacity),
			tokenRequest,
			utils.BoolString(s.Spec.RequiresRepublish),
			strings.Join(modes, ","),
			age,
		)

		fmt.Fprintln(w, csiDriverInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())

	return nil
}

// StorageClasses - a public function for searching storageclasses with keyword
func StorageClasses(ctx context.Context, opt *options.SearchOptions, keyword string) error {
	var storageClassInfo string

	storageClassList, err := utils.StorageClassList(ctx, opt)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.StorageClassesHeader)

	for _, s := range storageClassList.Items {
		// return all storages under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(s.Name, keyword)
			if match == opt.InvertMatch {
				continue
			}
		}

		age := utils.GetAge(time.Since(s.CreationTimestamp.Time))

		var isDefaultClass string
		for k, v := range s.Annotations {
			if k == "storageclass.kubernetes.io/is-default-class" && v == "true" {
				isDefaultClass = "(default)"
				break
			}
		}

		storageClassInfo = fmt.Sprintf(constants.StorageClassesRowTemplate,
			s.Name,
			isDefaultClass,
			s.Provisioner,
			*(s.ReclaimPolicy),
			*(s.VolumeBindingMode),
			utils.BoolString(s.AllowVolumeExpansion),
			age,
		)

		fmt.Fprintln(w, storageClassInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())

	return nil
}
