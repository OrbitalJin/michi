package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type Parser struct {
	config    *Config
	reExtract *regexp.Regexp
	reRemove  *regexp.Regexp
}

func NewParser(config *Config) (*Parser, error) {
	if config == nil {
		return nil, fmt.Errorf("Parser config.cannot be nil.")
	}

	reExtract, err := regexp.Compile(config.detectionPattern)

	if err != nil {
		fmt.Println(fmt.Errorf("failed to compile the extraction pattern: %w", err))
		return nil, err
	}

	reRemove, err := regexp.Compile(config.removalPattern)

	if err != nil {
		fmt.Println(fmt.Errorf("failed to compile the removal pattern: %w", err))
		return nil, err
	}

	return &Parser{
		config:    config,
		reExtract: reExtract,
		reRemove:  reRemove,
	}, nil
}

func (p Parser) Collect(input string) (*Result, error) {
	var result Result

	var words []string

	matches := p.reExtract.FindAllStringSubmatch(input, -1)

	if len(matches) >= 1 {
		for _, match := range matches {
			if len(match) > 1 {
				words = append(words, match[1])
			}
		}
	} else {
		words = make([]string, 0)
	}	

	query := p.reRemove.ReplaceAllString(input, "")
	query = strings.Join(strings.Fields(query), " ")
	query = strings.TrimSpace(query)

	result.Matches = words
	result.Query = query

	return &result, nil
}
