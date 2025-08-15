package lifecycle

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/server"
	"github.com/urfave/cli/v2"
)

func Doctor(server *server.Server) *cli.Command {
	return &cli.Command{
		Name:  "doctor",
		Usage: "check michi status",
		Action: func(ctx *cli.Context) error {
			pidFile := server.GetConfig().PidFile

			pid, err := readPidFile(pidFile)
			running := err == nil && processExists(pid)

			fmt.Printf("Running: %t\n", running)
			if running {
				fmt.Printf("PID: %d\n", pid)
			}
			return nil
		},
	}
}
