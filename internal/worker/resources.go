package worker

import (
	"context"
	"fmt"
	"os"

	"github.com/amirhnajafiz/sanjab/pkg/enum"

	"gopkg.in/yaml.v3"
	v13 "k8s.io/api/apps/v1"
	v14 "k8s.io/api/autoscaling/v1"
	"k8s.io/api/batch/v1beta1"
	v12 "k8s.io/api/core/v1"
	v15 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

const (
	LocalDir = "./local/tmp"
)

// master manages the workers of each resource
type master struct {
	Cfg         Config
	CephDisable bool
	Metrics     Metrics
}

// create a yaml file from our object
func (m master) createLocalFile(obj runtime.Object, path string) error {
	f, err := os.Create(path)
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

// export yaml file into ceph cluster
func (m master) exportYaml(obj runtime.Object, name string, path string) error {
	if err := m.createLocalFile(obj, path); err != nil {
		return fmt.Errorf("failed to create local file: %v", err)
	}

	if !m.CephDisable {
		if err := m.Cfg.Storage.Upload(name, path); err != nil {
			return fmt.Errorf("failed to upload file: %v", err)
		}
	}

	return nil
}

// ------------ resource methods -------------

func (m master) newPodResource() *worker {
	wo := newWorker(enum.PodResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.CoreV1().Pods(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v12.Pod)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newDeploymentResource() *worker {
	wo := newWorker(enum.DeploymentResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.AppsV1().Deployments(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v13.Deployment)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newServiceResource() *worker {
	wo := newWorker(enum.ServiceResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.CoreV1().Services(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v12.Service)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newCronjobResource() *worker {
	wo := newWorker(enum.CronjobResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.BatchV1().CronJobs(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v1beta1.CronJob)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newConfigmapResource() *worker {
	wo := newWorker(enum.ConfigMapResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.CoreV1().ConfigMaps(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v12.ConfigMap)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newSecretResource() *worker {
	wo := newWorker(enum.SecretResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.CoreV1().Secrets(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v12.Secret)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newServiceAccountResource() *worker {
	wo := newWorker(enum.ServiceAccountResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.CoreV1().ServiceAccounts(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v12.ServiceAccount)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newStatefulResource() *worker {
	wo := newWorker(enum.StatefulResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.AppsV1().StatefulSets(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v13.StatefulSet)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newHPAResource() *worker {
	wo := newWorker(enum.HPAResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.AutoscalingV1().HorizontalPodAutoscalers(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v14.HorizontalPodAutoscaler)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newIngressResource() *worker {
	wo := newWorker(enum.IngressResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.NetworkingV1().Ingresses(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v15.Ingress)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}

func (m master) newPVCResource() *worker {
	wo := newWorker(enum.PVCResource)

	if m.Cfg.Has(wo.Resource) {
		wo.Status = enum.PendingStatus
		wo.WatcherFunc = func(options v1.ListOptions) (watch.Interface, error) {
			m.Metrics.Observe(wo.GetResource())

			timeOut := int64(m.Cfg.Timeout)

			return m.Cfg.Client.CoreV1().PersistentVolumeClaims(m.Cfg.Namespace).Watch(context.Background(), v1.ListOptions{TimeoutSeconds: &timeOut})
		}
		wo.CallBack = func(event watch.Event) error {
			obj := event.Object.(*v12.PersistentVolumeClaim)
			name := obj.GetName()
			path := fmt.Sprintf("%s/%s.yaml", LocalDir, name)

			return m.exportYaml(obj, name, path)
		}
	}

	return wo
}
