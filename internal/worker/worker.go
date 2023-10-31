package worker

import "github.com/amirhnajafiz/sanjab/pkg/enum"

// Worker manages a resource by watching it
// if a resource is added, it will store in
// to storage
type Worker interface {
	Watch() error
	GetStatus() string
	GetResource() string
}

// Register system workers
func Register() []Worker {
	var workers []Worker

	return workers
}

type worker struct {
	CallBack func() error
	Status   enum.Status
	Resource enum.Resource
}

func (w worker) Watch() error {
	return w.CallBack()
}

func (w worker) GetStatus() string {
	return w.Status.ToString()
}

func (w worker) GetResource() string {
	return w.Resource.ToString()
}
