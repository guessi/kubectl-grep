package utils

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// DaemonSetList - return a list of DaemonSet(s)
func DaemonSetList(ctx context.Context, opt *options.SearchOptions) (*appsv1.DaemonSetList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().DaemonSets(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list DaemonSets: %w", err)
	}
	return list, nil
}

// DeploymentList - return a list of Deployment(s)
func DeploymentList(ctx context.Context, opt *options.SearchOptions) (*appsv1.DeploymentList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().Deployments(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list Deployments: %w", err)
	}
	return list, nil
}

// ReplicaSetList - return a list of ReplicaSet(s)
func ReplicaSetList(ctx context.Context, opt *options.SearchOptions) (*appsv1.ReplicaSetList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().ReplicaSets(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list ReplicaSets: %w", err)
	}
	return list, nil
}

// StatefulSetList - return a list of StatefulSet(s)
func StatefulSetList(ctx context.Context, opt *options.SearchOptions) (*appsv1.StatefulSetList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().StatefulSets(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list StatefulSets: %w", err)
	}
	return list, nil
}
