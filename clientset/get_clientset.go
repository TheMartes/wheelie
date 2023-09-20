package clientset

import (
	"log"
	kubernetesconfig "wheelie/kubernetes_config"

	"k8s.io/client-go/kubernetes"
)

var Context string

func GetClientSet() *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(kubernetesconfig.GetConfig(Context))
	if err != nil {
		log.Print(err.Error())
	}

	return clientset
}

func ChangeContext(ctx string) {
	Context = ctx
}
