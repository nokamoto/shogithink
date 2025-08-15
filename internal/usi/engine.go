package usi

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type logger interface {
	Log(string, ...any)
}

type handlerFunc func(io.Writer, []string) error

// Engine represents a USI engine.
// Engine implements the basic processing for USI engines (receiving commands and responding).
// It manages handlers (functions for each command) for each engine type in a map,
// and realizes engine behavior compliant with the USI protocol.
type Engine struct {
	name    string
	author  string
	logger  logger
	stdin   io.Reader
	stdout  io.Writer
	handler map[string]handlerFunc
}

// Run starts the main loop of the USI engine.
// It reads commands from stdin, logs received input, and dispatches each command to the appropriate handler.
// The method responds to standard USI commands ("usi", "isready", "quit") and delegates other commands to registered handlers.
// The loop continues until the input stream is closed or a "quit" command is received.
func (e *Engine) Run() error {
	res := func(msg string, args ...any) {
		send(e.stdout, e.logger, msg, args...)
	}
	scanner := bufio.NewScanner(e.stdin)
	for scanner.Scan() {
		line := scanner.Text()
		e.logger.Log("recv: %s", line)

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		ins := fields[0]
		switch ins {
		case "usi":
			res("id name %s", e.name)
			res("id author %s", e.author)
			res("usiok")

		case "isready":
			res("readyok")

		case "quit":
			e.logger.Log("quit received, exiting")
			return nil

		default:
			fn, ok := e.handler[ins]
			if !ok {
				continue
			}
			if err := fn(e.stdout, fields[1:]); err != nil {
				e.logger.Log("error handling command %s: %v", ins, err)
				continue
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}
	e.logger.Log("input stream closed, exiting")
	return nil
}
