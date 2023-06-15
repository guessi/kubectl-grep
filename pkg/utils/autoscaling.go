package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	autoscalingv2 "k8s.io/api/autoscaling/v2"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// HpaList - return a list of HPA(s)
func HpaList(opt *options.SearchOptions) *autoscalingv2.HorizontalPodAutoscalerList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AutoscalingV2().HorizontalPodAutoscalers(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get HPA List")
	}
	return list
}
