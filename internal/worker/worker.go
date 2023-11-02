package worker

import (
	"fmt"
	"log"
	"os"

	"github.com/amirhnajafiz/sanjab/pkg/enum"

	"gopkg.in/yaml.v3"
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
	// create a new master
	m := master{
		Cfg: cfg,
	}

	return []Worker{
		m.newPodResource(),
		m.newDeploymentResource(),
		m.newServiceResource(),
		m.newCronjobResource(),
		m.newConfigmapResource(),
		m.newSecretResource(),
		m.newServiceAccountResource(),
		m.newStatefulResource(),
		m.newHPAResource(),
		m.newIngressResource(),
		m.newPVCResource(),
	}
}

// each worker calls a watcher function to monitor resources
type worker struct {
	WatcherFunc func(options v1.ListOptions) (watch.Interface, error)
	CallBack    func(event watch.Event) error
	Status      enum.Status
	Resource    enum.Resource
}

// Watch a resource
func (w worker) Watch() error {
	return func() error {
		// disabled worker
		if w.Status == enum.DisableStatus {
			return nil
		}

		watcher, _ := toolsWatch.NewRetryWatcher("1", &cache.ListWatch{WatchFunc: w.WatcherFunc})

		for event := range watcher.ResultChan() {
			if event.Type == watch.Added {
				name := event.Object.GetObjectKind().GroupVersionKind().String()

				f, err := os.Create(name)
				if err != nil {
					log.Println(fmt.Sprintf("[worker][%s] failed to create file: %v", w.Resource, err))

					continue
				}

				err = yaml.NewEncoder(f).Encode(event.Object)
				if err != nil {
					log.Println(fmt.Sprintf("[worker][%s] failed to encode object: %v", w.Resource, err))

					continue
				}

				err = f.Close()
				if err != nil {
					log.Println(fmt.Sprintf("[worker][%s] failed to close file: %v", w.Resource, err))

					continue
				}
			}
		}

		return nil
	}()
}

// GetStatus of a worker
func (w worker) GetStatus() string {
	return w.Status.ToString()
}

// GetResource name of a worker
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
