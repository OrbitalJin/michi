package lifecycle

import (
	"fmt"
	"os"
	"strconv"

	"github.com/OrbitalJin/michi/internal/server"
	"github.com/urfave/cli/v2"
)

const (
	GREEN  = "\033[32m"
	RED    = "\033[31m"
	YELLOW = "\033[33m"
	RESET  = "\033[0m"
)

func Serve(server *server.Server) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "serve michi",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "detach"},
		},
		Action: func(ctx *cli.Context) error {
			pidFile := server.GetConfig().PidFile
			logFile := server.GetConfig().LogFile

			// Check if already running
			if pid, err := readPidFile(pidFile); err == nil {
				if processExists(pid) {
					fmt.Printf("%s●%s Server already running (PID: %d)\n",
						YELLOW, RESET, pid)
					return nil
				}
				// PID file exists but process is gone
				fmt.Printf("%s●%s Removing stale PID file (PID: %d not running)\n",
					RED, RESET, pid)
				_ = os.Remove(pidFile)
			}

			// Detached mode
			if ctx.Bool("detach") {
				if err := daemon(logFile); err != nil {
					fmt.Printf("%s●%s Failed to start server in background: %v\n",
						RED, RESET, err)
					return err
				}
				fmt.Printf("%s●%s Server started in background (logs: %s)\n",
					GREEN, RESET, logFile)
				return nil
			}

			// Foreground mode
			if err := os.WriteFile(
				pidFile,
				[]byte(strconv.Itoa(os.Getpid())),
				0644,
			); err != nil {
				fmt.Printf(
					"%s●%s Failed to write PID file: %v\n",
					RED,
					RESET,
					err,
				)
				return err
			}
			defer os.Remove(pidFile)

			fmt.Printf(
				"%s●%s Server running in foreground (PID: %d)\n",
				GREEN,
				RESET,
				os.Getpid(),
			)

			return server.Serve()
		},
	}
}
