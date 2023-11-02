package config

import "github.com/amirhnajafiz/sanjab/internal/storage"

func Default() Config {
	return Config{
		Storage: storage.Config{
			Host:   "",
			Secret: "",
			Access: "",
			Bucket: "",
		},
		Resources: []string{},
		Namespace: "",
		Timeout:   10,
		Port:      8080,
	}
}
