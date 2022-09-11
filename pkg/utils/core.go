package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// ConfigMapList - return a list of ConfigMap(s)
func ConfigMapList(opt *options.SearchOptions) *corev1.ConfigMapList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().ConfigMaps(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get ConfigMap List")
	}
	return list
}

// NodeList - return a list of Node(s)
func NodeList(opt *options.SearchOptions) *corev1.NodeList {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.CoreV1().Nodes().List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Node List")
	}
	return list
}

// PodList - return a list of Pod(s)
func PodList(opt *options.SearchOptions) *corev1.PodList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Pods(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Pod List")
	}
	return list
}

// SecretList - return a list of Secret(s)
func SecretList(opt *options.SearchOptions) *corev1.SecretList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Secrets(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Secret List")
	}
	return list
}

// ServiceList - return a list of Services
func ServiceAccountList(opt *options.SearchOptions) *corev1.ServiceAccountList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().ServiceAccounts(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get ServiceAccount List")
	}
	return list
}

// ServiceList - return a list of Services
func ServiceList(opt *options.SearchOptions) *corev1.ServiceList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Services(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Service List")
	}
	return list
}
