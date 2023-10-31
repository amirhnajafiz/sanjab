package worker

import "github.com/amirhnajafiz/sanjab/pkg/enum"

type Config struct {
	Resources []string
}

func (c Config) Has(resource enum.Resource) bool {
	for _, item := range c.Resources {
		if item == resource.ToString() {
			return true
		}
	}

	return false
}
