package server

import (
	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/OrbitalJin/qmuxr/internal/store"
)

type Config struct {
	bangParserCfg     *parser.Config
	shortcutParserCfg *parser.Config
	seshParserCfg     *parser.Config
	storeCfg          *store.Config
	serviceCgf        *service.Config
}

func NewConfig(
	bpCfg,
	scpCfg,
	seshCfg *parser.Config,
	sCfg *store.Config,
	svcCfg *service.Config,
) *Config {

	return &Config{
		bangParserCfg:     bpCfg,
		shortcutParserCfg: scpCfg,
		seshParserCfg:     seshCfg,
		storeCfg:          sCfg,
		serviceCgf:        svcCfg,
	}
}
