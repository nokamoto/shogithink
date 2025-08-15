package usi

import (
	"io"
)

// NewBridge creates a new USI engine bridge.
//
// This bridge implements minimal USI command handling.
// The core logic is to return dummy responses for "go" commands and leave "position" processing unimplemented.
func NewBridge(logger logger, stdin io.Reader, stdout io.Writer) *Engine {
	return &Engine{
		name:    "usi-bridge",
		author:  "github.com/nokamoto/shogithink",
		logger:  logger,
		stdin:   stdin,
		stdout:  stdout,
		handler: newBridge(logger),
	}
}

func newBridge(logger logger) map[string]handlerFunc {
	return map[string]handlerFunc{
		"position": func(w io.Writer, args []string) error {
			logger.Log("not implemented: position command")
			return nil
		},
		"go": func(w io.Writer, _ []string) error {
			send(w, logger, "info score cp 0")
			send(w, logger, "bestmove resign")
			return nil
		},
	}
}
