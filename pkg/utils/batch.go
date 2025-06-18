package utils

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	batchv1 "k8s.io/api/batch/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// CronJobList - return a list of CronJob(s)
func CronJobList(ctx context.Context, opt *options.SearchOptions) (*batchv1.CronJobList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.BatchV1().CronJobs(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list CronJobs: %w", err)
	}
	return list, nil
}

// JobList - return a list of Job(s)
func JobList(ctx context.Context, opt *options.SearchOptions) (*batchv1.JobList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.BatchV1().Jobs(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list Jobs: %w", err)
	}
	return list, nil
}
