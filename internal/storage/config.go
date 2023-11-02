package storage

type Config struct {
	Host   string `koanf:"host"`
	Access string `koanf:"access"`
	Secret string `koanf:"secret"`
	Bucket string `koanf:"bucket"`
}
