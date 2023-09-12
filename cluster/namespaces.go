package cluster

import (
	"context"
	"wheelie/clientset"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterNamespaces []string

func GetNamespaces(ctx context.Context) ClusterNamespaces {
	clientset := clientset.GetClientSet()
	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	prettyNamespaces := make([]string, len(namespaces.Items))

	for _, namespace := range namespaces.Items {
		prettyNamespaces = append(prettyNamespaces, namespace.Name)
	}

	return prettyNamespaces
}
