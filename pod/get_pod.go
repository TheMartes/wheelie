package pod

import (
	"context"
	"wheelie/clientset"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodListMetadata struct {
	ShortName string
	LongName  string
	Status    v1.PodPhase
}

func GetPods(ctx context.Context, namespace string) []PodListMetadata {
	// namespaces := a.GetNamespaces()
	clientset := clientset.GetClientSet()

	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	formattedPods := make([]PodListMetadata, len(pods.Items))

	for _, pod := range pods.Items {
		formattedPods = append(formattedPods, PodListMetadata{pod.Spec.Containers[0].Name, pod.Name, pod.Status.Phase})
	}

	return formattedPods
}
