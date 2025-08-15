package lifecycle

import (
	"fmt"
	"os"
	"strconv"

	"github.com/OrbitalJin/michi/internal/server"
	"github.com/urfave/cli/v2"
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

			if pid, err := readPidFile(pidFile); err == nil {
				if processExists(pid) {
					fmt.Printf("Server already running (PID %d)\n", pid)
					return nil
				}
				os.Remove(pidFile)
			}

			if ctx.Bool("detach") {
				return daemon(logFile)
			}

			if err := os.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0644); err != nil {
				return fmt.Errorf("failed to write PID file: %w", err)
			}

			defer os.Remove(pidFile)

			return server.Serve()
		},
	}
}
