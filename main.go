package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amirhnajafiz/sanjab/internal/config"
	internal "github.com/amirhnajafiz/sanjab/internal/http"
	"github.com/amirhnajafiz/sanjab/internal/storage"
	"github.com/amirhnajafiz/sanjab/internal/worker"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// load service configs
	configs := config.Load(os.Getenv("SJ_CONFIG_PATH"))

	// cluster client configs
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientSet, _ := kubernetes.NewForConfig(cfg)

	// open storage connection
	stg, err := storage.NewConnection(configs.Storage)
	if err != nil {
		panic(err)
	}

	// create workers
	workers := worker.Register(
		worker.Config{
			Storage:   stg,
			Client:    clientSet,
			Timeout:   configs.Timeout,
			Namespace: configs.Namespace,
			Resources: configs.Resources,
		},
		configs.CephDisable,
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
