package worker

import (
	"context"
	"fmt"
	"os"

	"github.com/amirhnajafiz/sanjab/pkg/enum"

	"gopkg.in/yaml.v3"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

type master struct {
	Cfg Config
}

func (m master) createLocalFile(obj runtime.Object, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	if err = yaml.NewEncoder(f).Encode(obj); err != nil {
		return fmt.Errorf("failed to encode object: %v", err)
	}

	if err = f.Close(); err != nil {
		return fmt.Errorf("failed to close file: %v", err)
	}

	return nil
}

func (m master) newPodResource() *worker {
	wo := newWorker(enum.PodResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.CoreV1().Pods(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v12.Pod)

			path := fmt.Sprintf("%s.yaml", obj.GetName())

			if err := m.createLocalFile(obj, path); err != nil {
				return fmt.Errorf("[worker][%s] failed to create local file: %v", wo.Resource, err)
			}

			if err := m.Cfg.Storage.Upload(obj.GetName(), path); err != nil {
				return fmt.Errorf("[worker][%s] failed to upload file: %v", wo.Resource, err)
			}

			return nil
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

func (m master) newServiceResource() *worker {
	wo := newWorker(enum.ServiceResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.CoreV1().Services(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newCronjobResource() *worker {
	wo := newWorker(enum.CronjobResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.BatchV1().CronJobs(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newConfigmapResource() *worker {
	wo := newWorker(enum.ConfigMapResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.CoreV1().ConfigMaps(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newSecretResource() *worker {
	wo := newWorker(enum.SecretResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.CoreV1().Secrets(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newServiceAccountResource() *worker {
	wo := newWorker(enum.ServiceAccountResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.CoreV1().ServiceAccounts(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newStatefulResource() *worker {
	wo := newWorker(enum.StatefulResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.AppsV1().StatefulSets(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newHPAResource() *worker {
	wo := newWorker(enum.HPAResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.AutoscalingV1().HorizontalPodAutoscalers(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newIngressResource() *worker {
	wo := newWorker(enum.IngressResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.NetworkingV1().Ingresses(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}

func (m master) newPVCResource() *worker {
	wo := newWorker(enum.PVCResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			timeOut := int64(m.Cfg.Timeout)
			return m.Cfg.Client.CoreV1().PersistentVolumeClaims(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
	}

	return wo
}
