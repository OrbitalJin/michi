package parser

import (
	"strings"
)

type QueryParser struct {
	bangParser     *Parser
	shortcutParser *Parser
}

func NewQueryParser(bpCfg, scpCfg *Config) (*QueryParser, error) {
	bp, err := NewParser(bpCfg)

	if err != nil {
		return nil, err
	}

	scp, err := NewParser(scpCfg)

	if err != nil {
		return nil, err
	}

	return &QueryParser{
		bangParser:     bp,
		shortcutParser: scp,
	}, nil
}

func (qd *QueryParser) ParseAction(query string) *QueryAction {
	trimmed := strings.TrimSpace(query)
	result, _ := qd.bangParser.Collect(trimmed)

	// Bangs
	if len(result.Matches) != 0 {
		return &QueryAction{
			Type:     BANG,
			Result:   result,
			RawQuery: trimmed,
		}
	}

	result, _ = qd.shortcutParser.Collect(trimmed)

	// Shortcuts
	if len(result.Matches) != 0 {
		return &QueryAction{
			Type:     SHORTCUT,
			Result:   result,
			RawQuery: trimmed,
		}
	}

	// Regular search
	return &QueryAction{
		Type: BANG,
		Result: &Result{
			Query: trimmed,
		},
		RawQuery: trimmed,
	}
}

func (qd *QueryParser) BangParser() *Parser {
	return qd.bangParser
}

func (qd *QueryParser) ShortcutParser() *Parser {
	return qd.shortcutParser
}
