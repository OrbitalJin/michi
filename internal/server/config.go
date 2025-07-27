package server

import (
	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/OrbitalJin/qmuxr/internal/store"
)

type Config struct {
	bangParserCfg     *parser.Config
	shortcutParserCfg *parser.Config
	storeCfg          *store.Config
	serviceCgf        *service.Config
}

func NewConfig(
	bpCfg,
	scpCfg *parser.Config,
	sCfg *store.Config,
	svcCfg *service.Config,
) *Config {

	return &Config{
		bangParserCfg:     bpCfg,
		shortcutParserCfg: scpCfg,
		storeCfg:          sCfg,
		serviceCgf:        svcCfg,
	}
}
