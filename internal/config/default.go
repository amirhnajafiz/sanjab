package config

func Default() Config {
	return Config{
		Resources: []string{},
		Namespace: "",
		Timeout:   10,
		Port:      8080,
	}
}
