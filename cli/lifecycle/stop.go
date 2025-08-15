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
				return fmt.Errorf("server not running")
			}

			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				return err
			}

			os.Remove(pidFile)
			fmt.Println("Server stopped")
			return nil
		},
	}
}
