package utils

import (
	"strings"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/guessi/kubectl-search/pkg/client"
)

var (
	clientset *kubernetes.Clientset
)

func init() {
	clientset = client.InitClient()
}

// GetAge - return human readable time expression
func GetAge(duration float64) (int, string) {
	var age int
	var unit string

	if duration >= 86400 {
		age = int(duration) / 86400
		unit = "d"
	} else if duration > 3600 {
		age = int(duration) / 3600
		unit = "h"
	} else if duration > 60 {
		age = int(duration) / 60
		unit = "m"
	} else {
		age = int(duration)
		unit = "s"
	}

	return age, unit
}

// setOptions - set common options for clientset
func setOptions(namespace string, allNamespaces bool, selector, fieldSelector string) (string, *metav1.ListOptions) {
	if len(namespace) <= 0 {
		namespace = "default"
	}

	if allNamespaces {
		namespace = ""
	}

	listOptions := &metav1.ListOptions{
		LabelSelector: selector,
		FieldSelector: fieldSelector,
	}
	return namespace, listOptions
}

// DaemonsetList - return a list of DaemonSet(s)
func DaemonsetList(namespace string, allNamespaces bool, selector, fieldSelector string) *appsv1.DaemonSetList {
	ns, o := setOptions(namespace, allNamespaces, selector, fieldSelector)
	list, err := clientset.AppsV1().DaemonSets(ns).List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get DaemonSet List")
	}
	return list
}

// DeploymentList - return a list of Deployment(s)
func DeploymentList(namespace string, allNamespaces bool, selector, fieldSelector string) *appsv1.DeploymentList {
	ns, o := setOptions(namespace, allNamespaces, selector, fieldSelector)
	list, err := clientset.AppsV1().Deployments(ns).List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Deployment List")
	}
	return list
}

// HpaList - return a list of HPA(s)
func HpaList(namespace string, allNamespaces bool, selector, fieldSelector string) *autoscalingv1.HorizontalPodAutoscalerList {
	ns, o := setOptions(namespace, allNamespaces, selector, fieldSelector)
	list, err := clientset.AutoscalingV1().HorizontalPodAutoscalers(ns).List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get HPA List")
	}
	return list
}

// PodList - return a list of Pod(s)
func PodList(namespace string, allNamespaces bool, selector, fieldSelector string) *corev1.PodList {
	ns, o := setOptions(namespace, allNamespaces, selector, fieldSelector)
	list, err := clientset.CoreV1().Pods(ns).List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Pod List")
	}
	return list
}

// NodeList - return a list of Node(s)
func NodeList(selector, fieldSelector string) *corev1.NodeList {
	_, o := setOptions("", false, selector, fieldSelector)
	list, err := clientset.CoreV1().Nodes().List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Node List")
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
