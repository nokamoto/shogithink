package main

import (
	"bufio"
	"log/slog"
	"os"
	"os/exec"

	"github.com/nokamoto/shogithink/internal/boilerplate"
	"github.com/nokamoto/shogithink/internal/observer"
)

func main() {
	port, err := boilerplate.GetObserverPort(8081)
	if err != nil {
		slog.Error("failed to get observer port", "error", err)
		os.Exit(1)
	}

	logger, err := observer.New(port)
	if err != nil {
		slog.Error("failed to create observer", "error", err)
		os.Exit(1)
	}

	usiPath := os.Getenv("USI_ENGINE_PATH")
	if usiPath == "" {
		slog.Error("USI_ENGINE_PATH environment variable is not set")
		os.Exit(1)
	}

	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		slog.Error("failed to create pipe", "error", err)
		os.Exit(1)
	}
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		slog.Error("failed to create stdin pipe", "error", err)
		os.Exit(1)
	}
	_, stderrWriter, err := os.Pipe()
	if err != nil {
		slog.Error("failed to create stderr pipe", "error", err)
		os.Exit(1)
	}
	cmd := exec.Command(usiPath)
	cmd.Stdout = stdoutWriter
	cmd.Stderr = stderrWriter
	cmd.Stdin = stdinReader

	// Start the USI engine for analysis in the background
	if err := cmd.Start(); err != nil {
		slog.Error("failed to start USI engine", "error", err)
		os.Exit(1)
	}

	logger.Log("USI engine started: path=%s, pid=%d", usiPath, cmd.Process.Pid)

	// Send USI commands to the engine
	if _, err := stdinWriter.Write([]byte("usi\n")); err != nil {
		slog.Error("failed to write to USI engine stdin", "error", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(stdoutReader)
	for scanner.Scan() {
		line := scanner.Text()
		logger.Log("USI engine output: %s", line)
		slog.Info("USI engine output", "line", line)
	}
	if err := scanner.Err(); err != nil {
		slog.Error("error reading USI engine output", "error", err)
		os.Exit(1)
	}
}
