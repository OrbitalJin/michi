package cli

import (
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/urfave/cli/v2"
)

type Cli struct {
	services *service.Services
	app      *cli.App
}

func New(services *service.Services) *Cli {
	return &Cli{
		services: services,
		app: &cli.App{
			Name: "michi",
		},
	}
}
