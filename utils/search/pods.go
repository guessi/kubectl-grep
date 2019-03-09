package search

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	client "github.com/guessi/kubectl-search/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	podsFields = "NAMESPACE\tNAME\tREADY\tSTATUS\tRESTART\tAGE"
)

// Pods - a public function for searching pods with keyword
func Pods(namespace string, allNamespaces bool, selector, fieldSelector, keyword string) {
	clientset := client.InitClient()

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

	pods, err := clientset.CoreV1().Pods(namespace).List(*listOptions)
	if err != nil {
		panic(err.Error())
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, podsFields)
	for _, p := range pods.Items {
		// return all pods under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(p.Name, keyword)
			if !match {
				continue
			}
		}

		var containerCount int = len(p.Status.ContainerStatuses)
		var readyCount int32
		var restartCount int32

		for _, cs := range p.Status.ContainerStatuses {
			restartCount += cs.RestartCount
			if cs.Ready {
				readyCount++
			}
		}

		age, ageUnit := getAge(time.Since(p.CreationTimestamp.Time).Seconds())

		pInfo := fmt.Sprintf("%s\t%s\t%d/%d\t%s\t%d\t%d%s",
			p.Namespace,
			p.Name,
			readyCount, containerCount,
			p.Status.Phase,
			restartCount,
			age, ageUnit,
		)
		fmt.Fprintln(w, pInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
