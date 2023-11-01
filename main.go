package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amirhnajafiz/sanjab/internal/config"
	internal "github.com/amirhnajafiz/sanjab/internal/http"
	"github.com/amirhnajafiz/sanjab/internal/worker"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// load service configs
	configs := config.Load()

	// cluster client configs
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

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

	// create a handler
	h := internal.Handler{
		Workers: workers,
	}

	http.HandleFunc("/", h.Index)
	http.HandleFunc("/health", h.Health)

	// start http server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", configs.Port), nil); err != nil {
		panic(err)
	}
}
