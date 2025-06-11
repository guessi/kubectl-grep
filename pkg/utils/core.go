package utils

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// ConfigMapList - return a list of ConfigMap(s)
func ConfigMapList(ctx context.Context, opt *options.SearchOptions) (*corev1.ConfigMapList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().ConfigMaps(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get ConfigMap List")
		return nil, fmt.Errorf("failed to list ConfigMaps: %w", err)
	}
	return list, nil
}

// NodeList - return a list of Node(s)
func NodeList(ctx context.Context, opt *options.SearchOptions) (*corev1.NodeList, error) {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.CoreV1().Nodes().List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Node List")
		return nil, fmt.Errorf("failed to list Nodes: %w", err)
	}
	return list, nil
}

// PodList - return a list of Pod(s)
func PodList(ctx context.Context, opt *options.SearchOptions) (*corev1.PodList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Pods(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Pod List")
		return nil, fmt.Errorf("failed to list Pods: %w", err)
	}
	return list, nil
}

// SecretList - return a list of Secret(s)
func SecretList(ctx context.Context, opt *options.SearchOptions) (*corev1.SecretList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Secrets(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Secret List")
		return nil, fmt.Errorf("failed to list Secrets: %w", err)
	}
	return list, nil
}

// ServiceAccountList - return a list of ServiceAccount(s)
func ServiceAccountList(ctx context.Context, opt *options.SearchOptions) (*corev1.ServiceAccountList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().ServiceAccounts(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get ServiceAccount List")
		return nil, fmt.Errorf("failed to list ServiceAccounts: %w", err)
	}
	return list, nil
}

// ServiceList - return a list of Service(s)
func ServiceList(ctx context.Context, opt *options.SearchOptions) (*corev1.ServiceList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Services(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Service List")
		return nil, fmt.Errorf("failed to list Services: %w", err)
	}
	return list, nil
}
