package lifecycle

import (
	"fmt"
	"os"
	"syscall"

	"github.com/OrbitalJin/michi/internal/server"
	"github.com/urfave/cli/v2"
)

func Stop(server *server.Server) *cli.Command {
	return &cli.Command{
		Name:  "stop",
		Usage: "stop michi",
		Action: func(ctx *cli.Context) error {
			pidFile := server.GetConfig().PidFile

			pid, err := readPidFile(pidFile)
			if err != nil {
				fmt.Printf("%s●%s Server not running (no PID file)\n", RED, RESET)
				return nil
			}

			// Try to stop the process
			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				fmt.Printf(
					"%s●%s Failed to stop server (PID: %d): %v\n",
					RED,
					RESET,
					pid,
					err,
				)
				return err
			}

			// Remove PID file
			_ = os.Remove(pidFile)

			fmt.Printf("%s●%s Server stopped (PID: %d)\n", GREEN, RESET, pid)
			return nil
		},
	}
}
