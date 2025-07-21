package parser

import (
	"fmt"
	"regexp"
)

// Parser type
type Parser struct {
	config *Config
	re     *regexp.Regexp
}

// Constructor
func NewParser(config *Config) (*Parser, error) {
	if config == nil {
		return nil, fmt.Errorf("Parser config.cannot be nil.")
	}

	re, err := regexp.Compile(config.pattern)

	if err != nil {
		return nil, err
	}

	return &Parser{
		config: config,
		re:     re,
	}, nil
}
