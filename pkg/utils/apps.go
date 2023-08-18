package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// DaemonsetList - return a list of DaemonSet(s)
func DaemonsetList(opt *options.SearchOptions) *appsv1.DaemonSetList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().DaemonSets(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get DaemonSet List")
	}
	return list
}

// DeploymentList - return a list of Deployment(s)
func DeploymentList(opt *options.SearchOptions) *appsv1.DeploymentList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().Deployments(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Deployment List")
	}
	return list
}

// ReplicaSetList - return a list of ReplicaSets
func ReplicaSetList(opt *options.SearchOptions) *appsv1.ReplicaSetList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().ReplicaSets(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get ReplicaSet List")
	}
	return list
}

// StatefulSetList - return a list of StatefulSets
func StatefulSetList(opt *options.SearchOptions) *appsv1.StatefulSetList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().StatefulSets(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get StatefulSet List")
	}
	return list
}
