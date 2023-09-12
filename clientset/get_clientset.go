package clientset

import (
	kubernetesconfig "wheelie/kubernetes_config"

	"k8s.io/client-go/kubernetes"
)

var Context string

func GetClientSet() *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(kubernetesconfig.GetConfig(Context))
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func ChangeContext(ctx string) {
	Context = ctx
}
