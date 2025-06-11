package utils

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	autoscalingv1 "k8s.io/api/autoscaling/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// HpaList - return a list of HorizontalPodAutoscaler(s)
func HpaList(ctx context.Context, opt *options.SearchOptions) (*autoscalingv1.HorizontalPodAutoscalerList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AutoscalingV1().HorizontalPodAutoscalers(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get HorizontalPodAutoscaler List")
		return nil, fmt.Errorf("failed to list HorizontalPodAutoscalers: %w", err)
	}
	return list, nil
}
