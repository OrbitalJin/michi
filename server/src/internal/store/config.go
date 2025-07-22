package store

type Config struct {
	path string
}

func NewConfig(path string) *Config {
	return &Config{
		path: path,
	}
}
