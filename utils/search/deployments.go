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
	deploymentsFields     = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tAGE"
	deploymentsFieldsWide = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tAGE\tCONTAINERS\tIMAGES"
	dInfo                 string
)

// Deployments - a public function for searching deployments with keyword
func Deployments(namespace string, allNamespaces bool, selector, fieldSelector, keyword string, wide bool) {
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

	deployments, err := clientset.AppsV1().Deployments(namespace).List(*listOptions)
	if err != nil {
		panic(err.Error())
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	if wide {
		fmt.Fprintln(w, deploymentsFieldsWide)
	} else {
		fmt.Fprintln(w, deploymentsFields)
	}
	for _, d := range deployments.Items {
		// return all deployments under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(d.Name, keyword)
			if !match {
				continue
			}
		}

		age, ageUnit := getAge(time.Since(d.CreationTimestamp.Time).Seconds())
		containers := d.Spec.Template.Spec.Containers

		if wide {
			names := []string{}
			images := []string{}

			for _, n := range containers {
				names = append(names, n.Name)
				images = append(images, n.Image)
			}

			dInfo = fmt.Sprintf("%s\t%s\t%d\t%d\t%d\t%d\t%d%s\t%s\t%s",
				d.Namespace,
				d.Name,
				d.Status.Replicas,
				d.Status.ReadyReplicas,
				d.Status.UpdatedReplicas,
				d.Status.AvailableReplicas,
				age, ageUnit,
				strings.Join(names, ","),
				strings.Join(images, ","),
			)
		} else {
			dInfo = fmt.Sprintf("%s\t%s\t%d\t%d\t%d\t%d\t%d%s",
				d.Namespace,
				d.Name,
				d.Status.Replicas,
				d.Status.ReadyReplicas,
				d.Status.UpdatedReplicas,
				d.Status.AvailableReplicas,
				age, ageUnit,
			)
		}
		fmt.Fprintln(w, dInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
