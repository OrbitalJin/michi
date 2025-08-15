package cli

import (
	"github.com/OrbitalJin/michi/cli/history"
	"github.com/OrbitalJin/michi/cli/lifecycle"
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
		Usage:                "A super-charger search engine multiplexer 🚀",
		EnableBashCompletion: true,
		Commands: []*v2.Command{
			history.Root(server.GetServices().GetHistoryService()),
			lifecycle.Serve(server),
			lifecycle.Stop(server),
			lifecycle.Doctor(server),
		},
	}
}
