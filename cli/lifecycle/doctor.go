package lifecycle

import (
	"fmt"
	"os"

	"github.com/OrbitalJin/michi/internal/server"
	"github.com/urfave/cli/v2"
)

func Doctor(server *server.Server) *cli.Command {
	return &cli.Command{
		Name:  "doctor",
		Usage: "check michi status",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "fix",
				Usage: "remove stale PID file",
				Value: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			fix := ctx.Bool("fix")

			pidFile := server.GetConfig().PidFile

			pid, err := readPidFile(pidFile)
			running := err == nil && processExists(pid)

			if running {
				fmt.Printf("%s●%s Running   (PID: %d)\n", GREEN, RESET, pid)
			} else if err == nil {
				fmt.Printf("%s●%s Stale PID file found (PID: %d not running)\n", YELLOW, RESET, pid)
				if fix {
					fmt.Printf("%s●%s Removing stale PID file (PID: %d not running)\n",
						RED, RESET, pid)
					_ = os.Remove(pidFile)
					fmt.Printf("%s●%s Michi should be ready to run\n", GREEN, RESET)
				}
			} else {
				fmt.Printf("%s●%s Not running\n", RED, RESET)
			}

			return nil
		},
	}
}
