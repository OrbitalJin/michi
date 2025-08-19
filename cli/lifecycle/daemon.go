package lifecycle

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

// daemon runs the server in the background and redirects logs to a file
func daemon(logFile string) error {
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable: %w", err)
	}

	log, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	cmd := exec.Command(exe, "serve")
	cmd.Stdout = log
	cmd.Stderr = log
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	// Start the process
	if err := cmd.Start(); err != nil {
		log.Close() // cleanup if start fails
		return fmt.Errorf("failed to start process: %w", err)
	}

	log.Close()

	fmt.Printf("Server started in background (PID %d)\n", cmd.Process.Pid)
	return nil
}

// readPidFile returns the pid of the server
func readPidFile(pidFile string) (int, error) {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, fmt.Errorf("failed to read pid file: %w", err)
	}
	return strconv.Atoi(string(data))
}

func processExists(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// Signal 0 is a way to check if the process exists without killing it
	return process.Signal(syscall.Signal(0)) == nil
}
