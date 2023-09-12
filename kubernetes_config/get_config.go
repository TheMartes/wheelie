package kubernetesconfig

import (
	"os"
	"path"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var KUBE_CONFIG = "/.kube/config"

type KubeContexts []string

func getKubeConfig() string {
	userHomeDirPath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return path.Join(userHomeDirPath, KUBE_CONFIG)
}

func GetConfig(ctx string) *rest.Config {
	kubeconfig := getKubeConfig()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{
			CurrentContext: ctx,
		}).ClientConfig()

	if err != nil {
		panic(err.Error())
	}

	return config
}

func GetKubeContexts() KubeContexts {
	kubeconfig := getKubeConfig()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{
			CurrentContext: "",
		}).RawConfig()

	if err != nil {
		panic(err)
	}

	kctx := make(KubeContexts, len(config.Contexts))

	for ctxName := range config.Contexts {
		kctx = append(kctx, ctxName)
	}

	return kctx
}
