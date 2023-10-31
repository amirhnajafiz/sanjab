package worker

import "github.com/amirhnajafiz/sanjab/pkg/enum"

// Worker manages a resource by watching it
// if a resource is added, it will store in
// to storage
type Worker interface {
	Watch() error
	State() enum.Status
}

// Register system workers
func Register() []Worker {
	var workers []Worker

	return workers
}
