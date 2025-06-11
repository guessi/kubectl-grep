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

// CronJobs - a public function for searching cronjobs with keyword
func CronJobs(ctx context.Context, opt *options.SearchOptions, keyword string) error {
	var cronjobInfo string

	cronjobList, err := utils.CronJobList(ctx, opt)
	if err != nil {
		return err
	}

	if len(cronjobList.Items) == 0 {
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

	fmt.Fprintln(w, constants.CronJobsHeader)

	for _, j := range cronjobList.Items {
		if !utils.MatchesKeyword(j.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(j.Name, opt.ExcludePattern) {
			continue
		}

		var lastScheduleTime string = "<none>"
		if j.Status.LastScheduleTime != nil {
			lastScheduleTime = utils.GetAge(time.Since(j.Status.LastScheduleTime.Time))
		}

		cronjobInfo = fmt.Sprintf(constants.CronJobsRowTemplate,
			j.Namespace,
			j.Name,
			j.Spec.Schedule,
			utils.BoolString(j.Spec.Suspend),
			len(j.Status.Active),
			lastScheduleTime,
			utils.GetAge(time.Since(j.CreationTimestamp.Time)),
		)
		fmt.Fprintln(w, cronjobInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())

	return nil
}
