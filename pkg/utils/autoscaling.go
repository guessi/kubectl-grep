package utils

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	autoscalingv2 "k8s.io/api/autoscaling/v2"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// HpaList - return a list of HorizontalPodAutoscaler(s)
func HpaList(ctx context.Context, opt *options.SearchOptions) (*autoscalingv2.HorizontalPodAutoscalerList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AutoscalingV2().HorizontalPodAutoscalers(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list HorizontalPodAutoscalers: %w", err)
	}
	return list, nil
}
