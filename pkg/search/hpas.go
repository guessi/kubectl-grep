package search

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	client "github.com/guessi/kubectl-search/pkg/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	hpasFields = "NAMESPACE\tNAME\tREFERENCE\tTARGETS\tMINPODS\tMAXPODS\tREPLICAS\tAGE"
)

// Hpas - a public function for searching hpas with keyword
func Hpas(namespace string, allNamespaces bool, selector, filedSelector, keyword string) {
	clientset := client.InitClient()

	if len(namespace) <= 0 {
		namespace = "default"
	}

	if allNamespaces {
		namespace = ""
	}

	listOptions := &metav1.ListOptions{
		LabelSelector: selector,
		FieldSelector: filedSelector,
	}

	hpas, err := clientset.AutoscalingV1().HorizontalPodAutoscalers(namespace).List(*listOptions)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, hpasFields)
	for _, h := range hpas.Items {
		// return all hpas under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(h.Name, keyword)
			if !match {
				continue
			}
		}

		age, ageUnit := getAge(time.Since(h.CreationTimestamp.Time).Seconds())

		hInfo := fmt.Sprintf("%s\t%s\t%s/%s\t%d%%/%d%%\t%d\t%d\t%d\t%d%s",
			h.Namespace,
			h.Name,
			h.Spec.ScaleTargetRef.Kind,
			h.Spec.ScaleTargetRef.Name,
			*h.Status.CurrentCPUUtilizationPercentage,
			*h.Spec.TargetCPUUtilizationPercentage,
			*h.Spec.MinReplicas,
			h.Spec.MaxReplicas,
			h.Status.CurrentReplicas,
			age, ageUnit,
		)
		fmt.Fprintln(w, hInfo)
	}
	w.Flush()

	fmt.Printf("%s", buf.String())
}
