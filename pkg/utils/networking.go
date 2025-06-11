package utils

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	networkingv1 "k8s.io/api/networking/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// IngressList - return a list of Ingress(es)
func IngressList(ctx context.Context, opt *options.SearchOptions) (*networkingv1.IngressList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.NetworkingV1().Ingresses(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Ingress List")
		return nil, fmt.Errorf("failed to list Ingresses: %w", err)
	}
	return list, nil
}
