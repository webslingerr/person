package config

type Config struct {
	PersonFileName string
	Path string
}

func Load() Config {
	cfg := Config{}

	cfg.Path = "./data"
	cfg.PersonFileName = "/data.json"

	return cfg
}