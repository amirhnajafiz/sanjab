package worker

import (
	"context"
	"log"

	"github.com/amirhnajafiz/sanjab/pkg/enum"

	cv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	toolsWatch "k8s.io/client-go/tools/watch"
)

func newPodResource(client *kubernetes.Clientset, cfg Config) *worker {
	w := newWorker(enum.PodResource)

	if cfg.Has(w.Resource) {
		w.Status = enum.PendingStatus
		w.CallBack = func() error {
			watchFunc := func(options v1.ListOptions) (watch.Interface, error) {
				timeOut := int64(60)

				return client.CoreV1().Pods("").Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
			}

			watcher, _ := toolsWatch.NewRetryWatcher("1", &cache.ListWatch{WatchFunc: watchFunc})

			for event := range watcher.ResultChan() {
				item := event.Object.(*cv1.Pod)

				if event.Type == watch.Added {
					// save item to storage
					log.Println(item)
				}
			}

			return nil
		}
	}

	return w
}

func newDeploymentResource(client *kubernetes.Clientset, cfg Config) *worker {
	w := newWorker(enum.DeploymentResource)

	if cfg.Has(w.Resource) {
		w.Status = enum.PendingStatus
		w.CallBack = func() error {
			// some logic

			return nil
		}
	}

	return w
}
