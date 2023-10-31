package worker

import (
	"context"

	"github.com/amirhnajafiz/sanjab/pkg/enum"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (w worker) newPodResource() *worker {
	wo := newWorker(enum.PodResource)

	if w.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(60)

			return w.Cfg.Client.CoreV1().Pods("").Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (w worker) newDeploymentResource() *worker {
	wo := newWorker(enum.DeploymentResource)

	if w.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(60)

			return w.Cfg.Client.AppsV1().Deployments("").Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}
