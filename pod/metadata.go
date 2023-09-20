package pod

import (
	"context"
	"log"
	"wheelie/clientset"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodInfo struct {
	Retrieved    bool
	Name         string
	Namespace    string
	StartTime    metav1.Time
	Status       corev1.PodPhase
	ImageName    string
	EnvVariables []corev1.EnvVar
	ReleaseName  string
}

func GetMetadata(ctx context.Context, name string, namespace string) PodInfo {
	clientset := clientset.GetClientSet()

	pod, err := clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Println("error loading pod info: ", err, "name: ", name, "namespace: ", namespace)
		return PodInfo{Retrieved: false}
	}

	return PodInfo{
		Retrieved:    true,
		Name:         pod.Name,
		Namespace:    pod.Namespace,
		StartTime:    *pod.Status.StartTime,
		Status:       pod.Status.Phase,
		ImageName:    pod.Spec.Containers[0].Image,
		EnvVariables: pod.Spec.Containers[0].Env,
		ReleaseName:  pod.Labels["app.kubernetes.io/instance"],
	}
}
