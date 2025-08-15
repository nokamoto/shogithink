package usi

import (
	"fmt"
	"io"
)

func send(stdout io.Writer, logger logger, msg string, args ...any) {
	s := fmt.Sprintf(msg, args...)
	fmt.Fprintln(stdout, s)
	logger.Log(s)
}
