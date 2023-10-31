package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amirhnajafiz/sanjab/internal/config"
	"github.com/amirhnajafiz/sanjab/internal/worker"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// cluster client configs
	cfg, _ := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	clientSet, _ := kubernetes.NewForConfig(cfg)

	// load service configs
	configs := config.Load()

	// create workers
	workers := worker.Register(clientSet, configs)

	// start workers
	for _, item := range workers {
		go func(w worker.Worker) {
			if err := w.Watch(); err != nil {
				log.Println(fmt.Errorf("[worker][%s] failed: %w", w.GetResource(), err))
			}
		}(item)
	}
}
