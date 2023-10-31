package worker

import "github.com/amirhnajafiz/sanjab/pkg/enum"

func newPodResource() *worker {
	w := &worker{
		Status:   enum.PendingStatus,
		Resource: enum.PodResource,
	}

	return w
}

func newDeploymentResource() *worker {
	w := &worker{
		Status:   enum.PendingStatus,
		Resource: enum.DeploymentResource,
	}

	return w
}
