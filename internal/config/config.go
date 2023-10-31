package config

type Config struct {
	Resources  []string `koanf:"resources"`
	KubeConfig string   `koanf:"kube_config"`
	Namespace  string   `koanf:"namespace"`
	Timeout    int      `koanf:"timeout"`
}

func Load() Config {
	return Default()
}
