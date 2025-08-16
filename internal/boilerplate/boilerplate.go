package boilerplate

import (
	"fmt"
	"os"
	"strconv"
)

// ErrIllegalPortRange is returned when the USI_OBSERVER_PORT is out of the valid range.
var ErrIllegalPortRange = fmt.Errorf("USI_OBSERVER_PORT must be between 1024 and 65535")

// GetObserverPort returns the port number for the observer service from the environment variable USI_OBSERVER_PORT.
// If the variable is not set, it returns the default port.
// It validates that the port is within the range of 1024 to 65535. If the port is out of range, it returns ErrIllegalPortRange.
func GetObserverPort(defaultPort int) (int, error) {
	portStr := os.Getenv("USI_OBSERVER_PORT")
	if portStr == "" {
		return defaultPort, nil
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, fmt.Errorf("invalid USI_OBSERVER_PORT: %w", err)
	}
	if port < 1024 || port > 65535 {
		return 0, fmt.Errorf("%w, got %d", ErrIllegalPortRange, port)
	}
	return port, nil
}
