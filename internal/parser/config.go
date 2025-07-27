package parser

import "fmt"

var core = `\b\w+\b`

type Config struct {
	prefix           string
	detectionPattern string
	removalPattern   string
}

func NewConfig(prefix string) *Config {
	return &Config{
		prefix:           prefix,
		detectionPattern: fmt.Sprintf(`%s(%s)`, prefix, core),
		removalPattern:   fmt.Sprintf(`%s%s`, prefix, core),
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
