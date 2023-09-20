package pod

import (
	"context"
	"log"
	"wheelie/clientset"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func GetDeployment(ctx context.Context, pod string, namespace string) string {
	cs := clientset.GetClientSet()
	podInfo, err := cs.CoreV1().Pods(namespace).Get(ctx, pod, v1.GetOptions{})
	if err != nil {
		log.Print(err) // TODO: Handle
	}

	releaseName := podInfo.Labels["app.kubernetes.io/instance"]

	d, err := cs.AppsV1().Deployments(namespace).Get(ctx, releaseName, v1.GetOptions{})
	if err != nil {
		log.Print(err)
	}
	y, err := yaml.Marshal(d)
	if err != nil {
		log.Print(err)
	}

	return string(y[:])
}
