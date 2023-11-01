package config

type Config struct {
	Resources []string `koanf:"resources"`
	Timeout   int      `koanf:"timeout"`
	Port      int      `koanf:"port"`
	Namespace string   `koanf:"namespace"`
}

func Load() Config {
	return Default()
}
