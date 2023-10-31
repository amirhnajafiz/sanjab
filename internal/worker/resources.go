package worker

import (
	"context"

	"github.com/amirhnajafiz/sanjab/pkg/enum"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func newPodResource() *worker {
	w := newWorker(enum.PodResource)

	if w.Cfg.Has(w.Resource) {
		w.Status = enum.PendingStatus
		w.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(60)

			return w.Cfg.Client.CoreV1().Pods("").Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return w
}

func newDeploymentResource() *worker {
	w := newWorker(enum.DeploymentResource)

	if w.Cfg.Has(w.Resource) {
		w.Status = enum.PendingStatus
		w.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(60)

			return w.Cfg.Client.AppsV1().Deployments("").Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return w
}
