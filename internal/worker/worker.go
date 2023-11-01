package worker

import (
	"log"

	"github.com/amirhnajafiz/sanjab/pkg/enum"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	toolsWatch "k8s.io/client-go/tools/watch"
)

// Worker manages a resource by watching it
// if a resource is added, it will store in
// to storage
type Worker interface {
	Watch() error
	GetStatus() string
	GetResource() string
}

// Register system workers
func Register(cfg Config) []Worker {
	var workers []Worker

	m := master{
		Cfg: cfg,
	}

	workers = append(workers, m.newPodResource())
	workers = append(workers, m.newDeploymentResource())

	return workers
}

type worker struct {
	WatcherFunc func(options v1.ListOptions) (watch.Interface, error)
	Status      enum.Status
	Resource    enum.Resource
}

func (w worker) Watch() error {
	return func() error {
		if w.Status == enum.DisableStatus {
			return nil
		}

		watcher, _ := toolsWatch.NewRetryWatcher("1", &cache.ListWatch{WatchFunc: w.WatcherFunc})

		for event := range watcher.ResultChan() {
			if event.Type == watch.Added {
				// save item to storage
				log.Println(event.Object.DeepCopyObject())
			}
		}

		return nil
	}()
}

func (w worker) GetStatus() string {
	return w.Status.ToString()
}

func (w worker) GetResource() string {
	return w.Resource.ToString()
}

// newWorker returns a raw worker with disabled status
func newWorker(resource enum.Resource) *worker {
	return &worker{
		Resource: resource,
		Status:   enum.DisableStatus,
	}
}
