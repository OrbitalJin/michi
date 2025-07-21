package parser

type Config struct {
	pattern string
	path string
}


func NewConfig(pattern string) *Config {
	return &Config{
		pattern: pattern,
	}
}

// Getter for string representation of regular expression pattern 
func (c *Config) GetPattern() string {
	return c.pattern
}

// Setter for string representation of regular expression pattern
func (c *Config) SetPattern(pattern string) bool {
	c.pattern = pattern
	return true
}

