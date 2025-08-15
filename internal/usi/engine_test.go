package usi

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testLogger struct{}

func (l *testLogger) Log(string, ...any) {}

type testcase struct {
	name    string
	input   []string
	want    []string
	wantErr error
}

func run(t *testing.T, cases []testcase, fn func(logger, io.Reader, io.Writer) *Engine) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var lines []string
			for _, line := range tt.input {
				lines = append(lines, line+"\n")
			}
			in := strings.NewReader(strings.Join(lines, ""))
			out := &bytes.Buffer{}
			logger := &testLogger{}
			// create the engine using the provided function
			e := fn(logger, in, out)
			err := e.Run()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("error = %v, want %v", err, tt.wantErr)
			}
			// test if output matches expected
			lines = []string{}
			for _, line := range tt.want {
				lines = append(lines, line+"\n")
			}
			if diff := cmp.Diff(strings.Join(lines, ""), out.String()); diff != "" {
				t.Errorf("output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEngine_Run(t *testing.T) {
	name := "test"
	author := "author"

	cases := []testcase{
		{
			name:  "USI information is returned if 'usi' is received",
			input: []string{"usi", "quit"},
			want: []string{
				fmt.Sprintf("id name %s", name),
				fmt.Sprintf("id author %s", author),
				"usiok",
			},
		},
		{
			name:  "'readyok' is returned if 'isready' is received",
			input: []string{"isready", "quit"},
			want:  []string{"readyok"},
		},
		{
			name:  "Nothing is returned if an unknown command is received",
			input: []string{"foo", "quit"},
		},
		{
			name:  "Nothing is returned if the handler returns an error",
			input: []string{"bar", "quit"},
		},
	}

	testEngine := func(logger logger, stdin io.Reader, stdout io.Writer) *Engine {
		return &Engine{
			name:   name,
			author: author,
			logger: logger,
			stdin:  stdin,
			stdout: stdout,
			handler: map[string]handlerFunc{
				"bar": func(w io.Writer, _ []string) error {
					return errors.New("failed")
				},
			},
		}
	}

	run(t, cases, testEngine)
}
