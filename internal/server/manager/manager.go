package manager

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/OrbitalJin/michi/internal"
	"github.com/OrbitalJin/michi/internal/server"
)

type ServerManager struct {
	server  *server.Server
	pidFile string
	logFile string
}

func NewServerManager(srv *server.Server) *ServerManager {
	return &ServerManager{
		server:  srv,
		pidFile: srv.GetConfig().PidFile,
		logFile: srv.GetConfig().LogFile,
	}
}

func (sm *ServerManager) Daemonize() error {
	if err := sm.canStart(); err != nil {
		return err
	}

	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable: %w", err)
	}

	logFile, err := os.OpenFile(sm.logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer logFile.Close()

	cmd := exec.Command(exe, "serve")
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	cmd.Env = os.Environ()
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid:     true,
		Foreground: false,
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start daemon: %w", err)
	}

	fmt.Printf("%s●%s Server started in background (PID: %d)\n", internal.ColorGreen, internal.ColorReset, cmd.Process.Pid)
	return nil
}

func (sm *ServerManager) RunForeground() error {
	if err := sm.canStart(); err != nil {
		return err
	}

	if err := sm.writePIDFile(); err != nil {
		return err
	}

	// Really important!!!
	defer os.Remove(sm.pidFile)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)
	go func() {
		fmt.Printf("%s●%s Server started in foreground\n", internal.ColorGreen, internal.ColorReset)
		fmt.Print("[>_] Press Ctrl+C to stop")
		errChan <- sm.server.Serve()
	}()

	select {
	case sig := <-sigChan:
		fmt.Printf("\n%s●%s Received signal %v, shutting down...\n", internal.ColorYellow, internal.ColorReset, sig)
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return sm.Shutdown()

	case err := <-errChan:
		return err
	}
}

func (sm *ServerManager) IsRunning() bool {
	err := sm.canStart()
	return err != nil
}

func (sm *ServerManager) ValiatePID() (bool, int) {
	_, err := os.Stat(sm.pidFile)

	if err != nil {
		return false, -1
	}

	data, err := os.ReadFile(sm.pidFile)

	if err != nil {
		return false, -1
	}

	pid, err := strconv.Atoi(string(data))

	if err != nil {
		return false, -1
	}

	return true, pid
}

func (sm *ServerManager) RemovePIDFile() error {
	if err := os.Remove(sm.pidFile); err != nil {
		return fmt.Errorf("Failed to remove PID file: %v\n", err)
	}
	return nil
}

func (sm *ServerManager) Shutdown() error {
	if err := sm.RemovePIDFile(); err != nil {
		return err
	}

	fmt.Printf("%s●%s Server stopped\n", internal.ColorGreen, internal.ColorReset)
	return nil
}

func (sm *ServerManager) GetServer() *server.Server {
	return sm.server
}

func (sm *ServerManager) ProcessExists(pid int) bool {
	proc, err := os.FindProcess(pid)

	if err != nil {
		return false
	}

	err = proc.Signal(syscall.Signal(0))

	if err == nil {
		return true
	}

	if err.Error() == "no such process" {
		return false
	}

	return false
}

func (sm *ServerManager) writePIDFile() error {
	pid := os.Getpid()
	if err := os.WriteFile(sm.pidFile, []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("Failed to write PID file: %w", err)
	}
	return nil
}

func (sm *ServerManager) canStart() error {
	ok, pid := sm.ValiatePID()

	if !ok {
		return nil
	}

	procExists := sm.ProcessExists(pid)

	if procExists {
		return fmt.Errorf("Server is already running (PID: %d)", pid)
	}

	return fmt.Errorf("Stale PID file found (PID: %d)", pid)
}
