package observer

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
)

type Observer struct {
	logs   []string
	logsMu sync.Mutex
	server *http.Server
}

// New creates a new Observer instance that listens on the specified port.
func New(port int) (*Observer, error) {
	o := &Observer{}
	o.server = &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			o.logsMu.Lock()
			defer o.logsMu.Unlock()
			for _, log := range o.logs {
				fmt.Fprintln(w, log)
			}
		}),
	}
	go func() {
		if err := o.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start HTTP server", "error", err)
		}
	}()
	return o, nil
}

// Stop gracefully stops the observer's HTTP server.
func (o *Observer) Stop() error {
	if o.server == nil {
		return nil
	}
	if err := o.server.Close(); err != nil {
		return fmt.Errorf("failed to stop server: %w", err)
	}
	return nil
}

// Log appends a formatted log message to the observer's logs.
func (o *Observer) Log(format string, args ...any) {
	o.logsMu.Lock()
	defer o.logsMu.Unlock()
	o.logs = append(o.logs, fmt.Sprintf(format, args...))
}
