package server

import (
	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/OrbitalJin/pow/internal/store"
)

type Config struct {
	parserCfg *parser.Config
	storeCfg  *store.Config
}

func NewConfig(pCfg *parser.Config, sCfg *store.Config) *Config {
	return &Config{
		parserCfg: pCfg,
		storeCfg:  sCfg,
	}
}
