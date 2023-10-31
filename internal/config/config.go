package config

import "github.com/amirhnajafiz/sanjab/internal/worker"

func Load() worker.Config {
	return Default()
}
