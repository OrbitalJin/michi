package cli

import (
	"github.com/OrbitalJin/michi/internal/server"
	v2 "github.com/urfave/cli/v2"
)

type Cli struct {
	server *server.Server
	cli    *v2.App
}

func New(server *server.Server) *v2.App {
	return &v2.App{
		Name:                 "michi",
		EnableBashCompletion: true,
		Commands: []*v2.Command{
			history(server.GetServices().GetHistoryService()),
			serve(server),
			stop(),
			doctor(),
		},
	}
}
