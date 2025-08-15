// usi-bridge is a simple USI bridge that listens to stdin and outputs to stdout.
// It also provides an HTTP server to view logs.
package main

import (
	"fmt"
	"os"

	"github.com/nokamoto/shogithink/internal/boilerplate"
	"github.com/nokamoto/shogithink/internal/observer"
	"github.com/nokamoto/shogithink/internal/usi"
)

func main() {
	port, err := boilerplate.GetObserverPort(8080)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting observer port: %v\n", err)
		os.Exit(1)
	}

	logger, err := observer.New(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating observer: %v\n", err)
		os.Exit(1)
	}

	bridge := usi.NewBridge(logger, os.Stdin, os.Stdout)
	if err := bridge.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
