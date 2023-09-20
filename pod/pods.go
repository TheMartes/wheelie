package pod

import (
	"context"
	"fmt"
	"log"
	"wheelie/clientset"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodListMetadata struct {
	ShortName   string
	LongName    string
	Status      v1.PodPhase
	ReleaseName string
}

func GetPods(ctx context.Context, namespace string) []PodListMetadata {
	// namespaces := a.GetNamespaces()
	clientset := clientset.GetClientSet()

	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}

	formattedPods := make([]PodListMetadata, len(pods.Items))

	for _, pod := range pods.Items {
		fmt.Print(pod.Labels["app.kubernetes.io/instance"])
		formattedPods = append(formattedPods, PodListMetadata{
			pod.Spec.Containers[0].Name,
			pod.Name,
			pod.Status.Phase,
			pod.Labels["app.kubernetes.io/instance"],
		})
	}

	return formattedPods
}
