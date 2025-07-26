package service

type Config struct {
	keepTrack       bool
	defaultProvider string
}

func NewConfig(keepTrack bool, dp string) *Config {
	return &Config{
		keepTrack:       keepTrack,
		defaultProvider: dp,
	}
}

func (c *Config) GetDefaultProvider() string {
	return c.defaultProvider
}

func (c *Config) ShouldKeepTrack() bool {
	return c.keepTrack
}
