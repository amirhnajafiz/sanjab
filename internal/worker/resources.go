package worker

import "github.com/amirhnajafiz/sanjab/pkg/enum"

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
