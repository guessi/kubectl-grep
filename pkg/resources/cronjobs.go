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

// CronJobs - a public function for searching cronjobs with keyword
func CronJobs(opt *options.SearchOptions, keyword string) {
	var cronjobInfo string

	cronjobList := utils.CronJobList(opt)

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.CronJobsHeader)

	for _, j := range cronjobList.Items {
		// return all jobs under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(j.Name, keyword)
			if !match {
				continue
			}
		}

		cronjobInfo = fmt.Sprintf(constants.CronJobsRowTemplate,
			j.Namespace,
			j.Name,
			j.Spec.Schedule,
			utils.BoolString(j.Spec.Suspend),
			len(j.Status.Active),
			utils.GetAge(time.Since(j.Status.LastScheduleTime.Time)),
			utils.GetAge(time.Since(j.CreationTimestamp.Time)),
		)
		fmt.Fprintln(w, cronjobInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
