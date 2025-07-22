package parser

type Config struct {
	detectionPattern string
	removalPattern   string
}

func NewConfig(detectionPattern, removalPattern string) *Config {
	return &Config{
		detectionPattern: detectionPattern,
		removalPattern:   removalPattern,
	}
}

func GetDefaultConfig() *Config {
	return &Config{
		detectionPattern: `!(\b\w+\b)`,
		removalPattern:   `!\b\w+\b`,
	}
}

func (c *Config) GetDetectionPattern() string {
	return c.detectionPattern
}

func (c *Config) GetRemovalPattern() string {
	return c.removalPattern
}
