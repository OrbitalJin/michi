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

func (c *Config) GetPattern() string {
	return c.detectionPattern
}

func (c *Config) SetPattern(pattern string) bool {
	c.detectionPattern = pattern
	return true
}
