package worker

import (
	"context"

	"github.com/amirhnajafiz/sanjab/pkg/enum"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type master struct {
	Cfg Config
}

func (m master) newPodResource() *worker {
	wo := newWorker(enum.PodResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.CoreV1().Pods(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newDeploymentResource() *worker {
	wo := newWorker(enum.DeploymentResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.AppsV1().Deployments(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}
