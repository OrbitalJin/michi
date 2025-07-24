package store

type Config struct {
	path string
	defaultProvider string
}

func NewConfig(path, defaultProvider string) *Config {
	return &Config{
		path: path,
		defaultProvider: defaultProvider,
	}
}

func (c *Config) GetPath() string {
	return c.path
}

func (c *Config) GetDefaultProvider() string {
	return c.defaultProvider
}
