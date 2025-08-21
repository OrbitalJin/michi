package lifecycle

import (
	"fmt"
	"syscall"

	"github.com/OrbitalJin/michi/internal"
	"github.com/OrbitalJin/michi/internal/server/manager"
	"github.com/urfave/cli/v2"
)

func Stop(sm *manager.ServerManager) *cli.Command {
    return &cli.Command{
        Name:  "stop",
        Usage: "stop michi",
        Action: func(ctx *cli.Context) error {
            ok, pid := sm.ValiatePID()
            
            if !ok {
                fmt.Printf("%s●%s Server not running (no PID file)\n", 
                    internal.ColorRed, internal.ColorReset)
                return nil
            }

            if !sm.ProcessExists(pid) {
                fmt.Printf("%s●%s Stale PID file found (PID: %d)\n", 
                    internal.ColorYellow, internal.ColorReset, pid)
                fmt.Printf("%s●%s Try running 'michi doctor --fix'\n", 
                    internal.ColorYellow, internal.ColorReset)
                return nil
            }

            // Try to stop the process
            if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
                fmt.Printf("%s●%s Failed to stop server (PID: %d): %v\n", 
                    internal.ColorRed, internal.ColorReset, pid, err)
                return nil
            }

            fmt.Printf("%s●%s Server stopped (PID: %d)\n", 
                internal.ColorGreen, internal.ColorReset, pid)
            return nil
        },
    }
}
