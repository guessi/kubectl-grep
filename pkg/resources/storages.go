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

// CsiDrivers - a public function for searching csidrivers with keyword
func CsiDrivers(opt *options.SearchOptions, keyword string) {
	var csiDriverInfo string

	csiDriverList := utils.CsiDriverList(opt)

	if len(csiDriverList.Items) <= 0 {
		if opt.AllNamespaces {
			fmt.Printf("No resources found.\n")
		} else {
			var ns = opt.Namespace
			if len(opt.Namespace) <= 0 {
				ns = "default"
			}
			fmt.Printf("No resources found in %s namespace.\n", ns)
		}
		return
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.CsiDriversHeader)

	for _, s := range csiDriverList.Items {
		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(s.Name, keyword)
			if !match {
				continue
			}
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
}

// StorageClasses - a public function for searching storageclasses with keyword
func StorageClasses(opt *options.SearchOptions, keyword string) {
	var storageClassInfo string

	storageClassList := utils.StorageClassList(opt)

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.StorageClassesHeader)

	for _, s := range storageClassList.Items {
		// return all storages under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(s.Name, keyword)
			if !match {
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
}
