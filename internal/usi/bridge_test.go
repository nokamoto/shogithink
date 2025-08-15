package usi

import (
	"testing"
)

func TestNewBridge(t *testing.T) {
	cases := []testcase{
		{
			name:  "not implemented: position command",
			input: []string{"position startpos", "quit"},
			want:  []string{},
		},
		{
			name:  "go command returns dummy responses",
			input: []string{"go", "quit"},
			want: []string{
				"info score cp 0",
				"bestmove resign",
			},
		},
	}

	run(t, cases, NewBridge)
}
