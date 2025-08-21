package lifecycle

import (
	"github.com/OrbitalJin/michi/internal/server/manager"
	"github.com/urfave/cli/v2"
)

func Serve(sm *manager.ServerManager) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "serve michi",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "detach"},
		},
		Action: func(ctx *cli.Context) error {

			if ctx.Bool("detach") {
				return sm.Daemonize()
			}

			return sm.RunForeground()
		},
	}
}
