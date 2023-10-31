package main

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/sanjab/internal/config"
	"github.com/amirhnajafiz/sanjab/internal/worker"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// load service configs
	configs := config.Load()

	// cluster client configs
	cfg, _ := clientcmd.BuildConfigFromFlags("", configs.KubeConfig)
	clientSet, _ := kubernetes.NewForConfig(cfg)

	// create workers
	workers := worker.Register(
		worker.Config{
			Client:    clientSet,
			Timeout:   configs.Timeout,
			Namespace: configs.Namespace,
			Resources: configs.Resources,
		},
	)

	// start workers
	for _, item := range workers {
		go func(w worker.Worker) {
			if err := w.Watch(); err != nil {
				log.Println(fmt.Errorf("[worker][%s] failed: %w", w.GetResource(), err))
			}
		}(item)
	}
}
