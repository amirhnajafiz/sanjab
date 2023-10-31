package config

func Default() Config {
	return Config{
		Resources:  []string{},
		Namespace:  "",
		KubeConfig: "",
		Timeout:    10,
	}
}
