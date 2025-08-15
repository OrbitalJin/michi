package lifecycle

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

// daemon runs the server in the background
func daemon(logFile string) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	cmd := exec.Command(exe, "serve")
	cmd.Stdout, _ = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	cmd.Stderr = cmd.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	if err := cmd.Start(); err != nil {
		return err
	}

	fmt.Printf("Server started in background (PID %d)\n", cmd.Process.Pid)
	return nil
}

// readPidFile returns the pid of the server
func readPidFile(pidFile string) (int, error) {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(data))
}

// processExists returns true if the process exists
func processExists(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	return process.Signal(syscall.Signal(0)) == nil
}
