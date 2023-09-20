package main

import (
	"context"
	"wheelie/clientset"
	"wheelie/cluster"
	kubernetesconfig "wheelie/kubernetes_config"
	"wheelie/pod"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetNamespaces() cluster.ClusterNamespaces {
	return cluster.GetNamespaces(a.ctx)
}

func (a *App) ChangeContext(ctx string) {
	clientset.ChangeContext(ctx)
}

func (a *App) GetPodsFromNamespace(namespace string) []pod.PodListMetadata {
	return pod.GetPods(a.ctx, namespace)
}

func (a *App) GetLogsFromPod(podName string, podNamespace string, podTailLines int) string {
	return pod.GetLogs(a.ctx, podName, podNamespace, podTailLines)
}

func (a *App) GetPodMetadata(podName string, podNamespace string) pod.PodInfo {
	return pod.GetMetadata(a.ctx, podName, podNamespace)
}

func (a *App) GetKubernetesContext() kubernetesconfig.KubeContexts {
	return kubernetesconfig.GetKubeContexts()
}

func (a *App) GetDeployment(podName string, podNamespace string) string {
	return pod.GetDeployment(a.ctx, podName, podNamespace)
}
