// usi-bridge is a simple USI bridge that listens to stdin and outputs to stdout.
// It also provides an HTTP server to view logs.
package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/nokamoto/shogithink/internal/usi"
)

var (
	logs   []string
	logsMu sync.Mutex
)

func appendLog(s string) {
	logsMu.Lock()
	defer logsMu.Unlock()
	logs = append(logs, s)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	logsMu.Lock()
	defer logsMu.Unlock()
	for _, log := range logs {
		fmt.Fprintln(w, log)
	}
}

func startHTTPServer() {
	http.HandleFunc("/", logsHandler)
	http.ListenAndServe(":8080", nil)
}

type httpServerLogger struct{}

func (httpServerLogger) Log(format string, args ...any) {
	appendLog(fmt.Sprintf(format, args...))
}

func main() {
	go startHTTPServer()
	logger := httpServerLogger{}
	bridge := usi.NewBridge(logger, os.Stdin, os.Stdout)
	if err := bridge.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
