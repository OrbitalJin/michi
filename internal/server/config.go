package server

import (
	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/OrbitalJin/qmuxr/internal/store"
)

type Config struct {
	parserCfg  *parser.Config
	storeCfg   *store.Config
	serviceCgf *service.Config
}

func NewConfig(pCfg *parser.Config, sCfg *store.Config, svcCfg *service.Config) *Config {
	return &Config{
		parserCfg:  pCfg,
		storeCfg:   sCfg,
		serviceCgf: svcCfg,
	}
}
