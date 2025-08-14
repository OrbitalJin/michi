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

var pidFile = "/tmp/michi.pid"

func serve(server *server.Server) *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "serve michi",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "detach",
			},
			&cli.BoolFlag{
				Name:   "child-process",
				Hidden: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			// Skip the isRunning check if this is a child process
			if !ctx.Bool("child-process") {
				running, pid, err := isRunning()
				if err != nil {
					return err
				}

				if running {
					fmt.Printf("Server is already running with PID %d\n", pid)
					return nil
				}
			}

			detach := ctx.Bool("detach")

			if detach {
				if err := serveInBackground(); err != nil {
					return fmt.Errorf("failed to serve detached: %w", err)
				}
				return nil
			}

			if err := server.Serve(); err != nil {
				return fmt.Errorf("failed to serve: %w", err)
			}

			return nil
		},
	}
}

func stop() *cli.Command {
	return &cli.Command{
		Name:  "stop",
		Usage: "stop michi",
		Action: func(ctx *cli.Context) error {

			running, pid, err := isRunning()
			if err != nil {
				return err
			}

			if !running {
				fmt.Println("Server is not running.")
				return nil
			}

			process, err := os.FindProcess(pid)
			if err != nil {
				return err
			}

			if err := process.Signal(syscall.SIGTERM); err != nil {
				return err
			}

			fmt.Println("Server stopped")
			os.Remove(pidFile)
			return nil
		},
	}
}

func doctor() *cli.Command {
	return &cli.Command{
		Name:  "doctor",
		Usage: "check michi",
		Action: func(ctx *cli.Context) error {
			running, pid, err := isRunning()
			if err != nil {
				return err
			}

			fmt.Printf("Running: %t\n", running)
			fmt.Printf("PID: %d\n", pid)
			return nil
		},
	}
}

func serveInBackground() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	// Create log files for the background process
	stdout, err := os.OpenFile("/tmp/michi.stdout.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	stderr, err := os.OpenFile("/tmp/michi.stderr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		stdout.Close()
		return err
	}

	// Start a new process without --detach flag
	// Add a special flag to indicate this is a child process
	cmd := exec.Command(exe, "serve", "--child-process")
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Stdin = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	if err := cmd.Start(); err != nil {
		stdout.Close()
		stderr.Close()
		return err
	}

	// Save PID
	if err := os.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0644); err != nil {
		return err
	}

	fmt.Printf("Server started in background with PID %d\n", cmd.Process.Pid)
	fmt.Printf("Logs available at /tmp/michi.stdout.log and /tmp/michi.stderr.log\n")

	return nil
}

func isRunning() (bool, int, error) {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false, 0, nil
		}
		return false, 0, err
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return false, 0, fmt.Errorf("invalid PID in file: %w", err)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return false, 0, nil
	}

	err = process.Signal(syscall.Signal(0))
	if err != nil {
		os.Remove(pidFile)
		return false, 0, nil
	}

	return true, pid, nil
}
