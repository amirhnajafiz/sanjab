package worker

import (
	"github.com/amirhnajafiz/sanjab/internal/storage"
	"github.com/amirhnajafiz/sanjab/pkg/enum"

	"k8s.io/client-go/kubernetes"
)

type Config struct {
	Storage   *storage.Storage
	Client    *kubernetes.Clientset
	Resources []string
	Timeout   int
	Namespace string
}

func (c Config) Has(resource enum.Resource) bool {
	for _, item := range c.Resources {
		if item == resource.ToString() {
			return true
		}
	}

	return false
}
