package pod

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"wheelie/clientset"

	corev1 "k8s.io/api/core/v1"
)

func GetLogs(ctx context.Context, podName string, podNamespace string, tl int) string {
	var podLogOpts corev1.PodLogOptions

	if tl == 0 {
		podLogOpts = corev1.PodLogOptions{}
	} else {
		tailLines := int64(tl)
		podLogOpts = corev1.PodLogOptions{
			TailLines: &tailLines,
		}
	}

	clientset := clientset.GetClientSet()
	req := clientset.CoreV1().Pods(podNamespace).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(ctx)
	if err != nil {
		return "error in opening stream"
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "error in copy information from podLogs to buf"
	}

	str := buf.String()

	return str
}

func SaveLogsFromPod(logs string) {
	f, err := os.Create("logs.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(logs)

	if err != nil {
		log.Fatal(err)
	}
}
