package utils

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/duration"
	"k8s.io/client-go/kubernetes"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

var (
	clientset *kubernetes.Clientset
)

func init() {
	clientset = client.InitClient()
}

// setOptions - set common options for clientset
func setOptions(opt *options.SearchOptions) (string, *metav1.ListOptions) {
	var namespace string
	if len(opt.Namespace) <= 0 {
		namespace = "default"
	} else {
		namespace = opt.Namespace
	}

	if opt.AllNamespaces {
		namespace = ""
	}

	listOptions := &metav1.ListOptions{
		LabelSelector: opt.Selector,
		FieldSelector: opt.FieldSelector,
	}
	return namespace, listOptions
}

// DaemonsetList - return a list of DaemonSet(s)
func DaemonsetList(opt *options.SearchOptions) *appsv1.DaemonSetList {
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().DaemonSets(ns).List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get DaemonSet List")
	}
	return list
}

// DeploymentList - return a list of Deployment(s)
func DeploymentList(opt *options.SearchOptions) *appsv1.DeploymentList {
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().Deployments(ns).List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Deployment List")
	}
	return list
}

// HpaList - return a list of HPA(s)
func HpaList(opt *options.SearchOptions) *autoscalingv1.HorizontalPodAutoscalerList {
	ns, o := setOptions(opt)
	list, err := clientset.AutoscalingV1().HorizontalPodAutoscalers(ns).List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get HPA List")
	}
	return list
}

// PodList - return a list of Pod(s)
func PodList(opt *options.SearchOptions) *corev1.PodList {
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Pods(ns).List(*o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Pod List")
	}
	return list
}

// NodeList - return a list of Node(s)
func NodeList(opt *options.SearchOptions) *corev1.NodeList {
	_, o := setOptions(opt)
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

// GetAge - return human readable time expression
func GetAge(d time.Duration) string {
	return duration.HumanDuration(d)
}
