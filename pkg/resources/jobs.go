package resources

import (
	"bytes"
	"fmt"
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

	if len(jobList.Items) == 0 {
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

	fmt.Fprintln(w, constants.JobsHeader)

	for _, j := range jobList.Items {
		var jobDuration string

		if !utils.MatchesKeyword(j.Name, keyword, opt.InvertMatch) {
			continue
		}

		if utils.ShouldExcludeResource(j.Name, opt.ExcludePattern) {
			continue
		}

		age := utils.GetAge(time.Since(j.CreationTimestamp.Time))

		completions := j.Spec.Completions
		succeeded := j.Status.Succeeded

		if succeeded > 0 && j.Status.StartTime != nil && j.Status.CompletionTime != nil {
			start := j.Status.StartTime.Time
			end := j.Status.CompletionTime.Time
			jobDuration = duration.HumanDuration(end.Sub(start))
		} else {
			jobDuration = age
		}

		var completionsValue int32 = 0
		if completions != nil {
			completionsValue = *completions
		}

		jobInfo = fmt.Sprintf(constants.JobsRowTemplate,
			j.Namespace,
			j.Name,
			succeeded, completionsValue,
			jobDuration,
			age,
		)
		fmt.Fprintln(w, jobInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
