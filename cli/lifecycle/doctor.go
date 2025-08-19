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

			if running {
				fmt.Printf("%s●%s Running   (PID: %d)\n", GREEN, RESET, pid)
			} else if err == nil {
				// PID file exists but process is gone
				fmt.Printf("%s●%s Stale PID file found (PID: %d not running)\n", YELLOW, RESET, pid)
			} else {
				// No PID file or unreadable
				fmt.Printf("%s●%s Not running\n", RED, RESET)
			}

			return nil
		},
	}
}
