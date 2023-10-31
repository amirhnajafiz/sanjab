package worker

import "github.com/amirhnajafiz/sanjab/pkg/enum"

func newPodResource(cfg Config) *worker {
	w := newWorker(enum.PodResource)

	if cfg.Has(w.Resource) {
		w.Status = enum.PendingStatus
		w.CallBack = func() error {
			// some logic

			return nil
		}
	}

	return w
}

func newDeploymentResource(cfg Config) *worker {
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
