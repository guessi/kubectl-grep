package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	networkingv1 "k8s.io/api/networking/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// IngressList - return a list of Ingresses
func IngressList(opt *options.SearchOptions) *networkingv1.IngressList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.NetworkingV1().Ingresses(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Ingress List")
	}
	return list
}
