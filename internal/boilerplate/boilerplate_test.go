package boilerplate

import (
	"errors"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetObserverPort(t *testing.T) {
	defaultPort := 8080

	type testcase struct {
		name    string
		env     string
		want    int
		wantErr error
	}

	testcases := []testcase{
		{
			name: "get default port if the environment variable is not set",
			want: defaultPort,
		},
		{
			name: "get valid port from environment variable",
			env:  "12345",
			want: 12345,
		},
		{
			name:    "ErrIllegalPortRange if port is below 1024",
			env:     "1023",
			wantErr: ErrIllegalPortRange,
		},
		{
			name:    "ErrIllegalPortRange if port is above 65535",
			env:     "65536",
			wantErr: ErrIllegalPortRange,
		},
		{
			name:    "error if environment variable is not a number",
			env:     "not_a_number",
			wantErr: strconv.ErrSyntax,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != "" {
				t.Setenv("USI_OBSERVER_PORT", tt.env)
			}

			got, err := GetObserverPort(defaultPort)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetObserverPort() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetObserverPort() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
