package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/OrbitalJin/michi/internal/server"
	"github.com/urfave/cli/v2"
)

const pidFile = "/tmp/michi.pid"

func serve(server *server.Server) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "serve michi",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "detach"},
		},
		Action: func(ctx *cli.Context) error {
			if pid, err := readPidFile(); err == nil {
				if processExists(pid) {
					fmt.Printf("Server already running (PID %d)\n", pid)
					return nil
				}
				os.Remove(pidFile)
			}

			if ctx.Bool("detach") {
				return daemonize()
			}

			if err := os.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0644); err != nil {
				return fmt.Errorf("failed to write PID file: %w", err)
			}
			defer os.Remove(pidFile)

			return server.Serve()
		},
	}
}

func stop() *cli.Command {
	return &cli.Command{
		Name:  "stop",
		Usage: "stop michi",
		Action: func(ctx *cli.Context) error {
			pid, err := readPidFile()
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

func doctor() *cli.Command {
	return &cli.Command{
		Name:  "doctor",
		Usage: "check michi status",
		Action: func(ctx *cli.Context) error {
			pid, err := readPidFile()
			running := err == nil && processExists(pid)

			fmt.Printf("Running: %t\n", running)
			if running {
				fmt.Printf("PID: %d\n", pid)
			}
			return nil
		},
	}
}

func daemonize() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	cmd := exec.Command(exe, "serve")
	cmd.Stdout, _ = os.OpenFile("/tmp/michi.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	cmd.Stderr = cmd.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	if err := cmd.Start(); err != nil {
		return err
	}

	fmt.Printf("Server started in background (PID %d)\n", cmd.Process.Pid)
	return nil
}

func readPidFile() (int, error) {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(data))
}

func processExists(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	return process.Signal(syscall.Signal(0)) == nil
}
