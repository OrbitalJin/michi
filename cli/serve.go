package cli

import (
	"github.com/OrbitalJin/michi/internal/server"
	"github.com/urfave/cli/v2"
)


func serve(server *server.Server) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "serve michi",
		Action: func(ctx *cli.Context) error {
			server.Serve()
			return nil
		},
	}
}
