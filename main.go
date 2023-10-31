package main

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	clientset, _ := kubernetes.NewForConfig(config)
}
