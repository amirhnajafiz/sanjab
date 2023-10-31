package worker

import "github.com/amirhnajafiz/sanjab/pkg/enum"

type Config struct {
	Resources []enum.Resource
}

func (c Config) Has(resource string) bool {
	for _, item := range c.Resources {
		if item.ToString() == resource {
			return true
		}
	}

	return false
}
