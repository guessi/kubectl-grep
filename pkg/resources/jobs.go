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
	"k8s.io/apimachinery/pkg/util/duration"
)

// Jobs - a public function for searching jobs with keyword
func Jobs(opt *options.SearchOptions, keyword string) {
	var jobInfo string

	jobList := utils.JobList(opt)

	if len(jobList.Items) <= 0 {
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

	fmt.Fprintln(w, constants.JobsHeader)

	for _, j := range jobList.Items {
		var jobDuration string

		// return all if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(j.Name, keyword)
			if match == opt.InvertMatch {
				continue
			}
		}

		age := utils.GetAge(time.Since(j.CreationTimestamp.Time))

		completions := j.Spec.Completions
		succeeded := j.Status.Succeeded

		if succeeded > 0 {
			start := j.Status.StartTime.Time
			end := j.Status.CompletionTime.Time
			jobDuration = duration.HumanDuration(end.Sub(start))
		} else {
			jobDuration = age
		}

		jobInfo = fmt.Sprintf(constants.JobsRowTemplate,
			j.Namespace,
			j.Name,
			succeeded, *completions,
			jobDuration,
			age,
		)
		fmt.Fprintln(w, jobInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
