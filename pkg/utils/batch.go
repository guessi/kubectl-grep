package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/batch/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// JobList - return a list of Jobs
func JobList(opt *options.SearchOptions) *v1.JobList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.BatchV1().Jobs(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Job List")
	}
	return list
}
