package server

import (
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/store"
)

type Config struct {
	bangParserCfg     *parser.Config
	shortcutParserCfg *parser.Config
	seshParserCfg     *parser.Config
	storeCfg          *store.Config
	serviceCgf        *service.Config
	port string
}

func NewConfig(
	port string,
	bpCfg,
	scpCfg,
	seshCfg *parser.Config,
	sCfg *store.Config,
	svcCfg *service.Config,
) *Config {

	return &Config{
		port:              port,
		bangParserCfg:     bpCfg,
		shortcutParserCfg: scpCfg,
		seshParserCfg:     seshCfg,
		storeCfg:          sCfg,
		serviceCgf:        svcCfg,
	}
}
