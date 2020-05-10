package utils

import (
	"context"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/duration"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// setOptions - set common options for clientset
func setOptions(opt *options.SearchOptions) (string, *metav1.ListOptions) {
	// set default namespace as "default"
	namespace := "default"

	// override `namespace` if `--all-namespaces` exist
	if opt.AllNamespaces {
		namespace = ""
	} else {
		if len(opt.Namespace) > 0 {
			namespace = opt.Namespace
		} else {
			ns, _, err := client.ClientConfig().Namespace()
			if err != nil {
				log.WithFields(log.Fields{
					"err": err.Error(),
				}).Debug("Failed to resolve namespace")
			} else {
				namespace = ns
			}
		}
	}

	// retrieve listOptions from meta
	listOptions := &metav1.ListOptions{
		LabelSelector: opt.Selector,
		FieldSelector: opt.FieldSelector,
	}
	return namespace, listOptions
}

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

// HpaList - return a list of HPA(s)
func HpaList(opt *options.SearchOptions) *autoscalingv1.HorizontalPodAutoscalerList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AutoscalingV1().HorizontalPodAutoscalers(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get HPA List")
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

// StatefulSetList - return a list of StatefulSets
func StatefulSetList(opt *options.SearchOptions) *appsv1.StatefulSetList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().StatefulSets(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get .StatefulSet List")
	}
	return list
}

// TrimQuoteAndSpace - remove Spaces, Tabs, SingleQuotes, DoubleQuites
func TrimQuoteAndSpace(input string) string {
	if len(input) >= 2 {
		if input[0] == '"' && input[len(input)-1] == '"' {
			return input[1 : len(input)-1]
		}
		if input[0] == '\'' && input[len(input)-1] == '\'' {
			return input[1 : len(input)-1]
		}
	}
	return strings.TrimSpace(input)
}

// GetAge - return human readable time expression
func GetAge(d time.Duration) string {
	return duration.HumanDuration(d)
}
