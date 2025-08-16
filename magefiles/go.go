package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Go provides Mage targets for Go-related tasks.
type Go mg.Namespace

// Tidy runs 'go mod tidy' to clean up the go.mod file.
func (Go) Tidy() error {
	return sh.Run("go", "mod", "tidy")
}

// Install installs Go commands.
func (Go) Install() error {
	commands := []string{
		"golang.org/x/tools/cmd/goimports@latest",
		"mvdan.cc/gofumpt@latest",
		"github.com/google/yamlfmt/cmd/yamlfmt@latest",
		"github.com/bufbuild/buf/cmd/buf@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"connectrpc.com/connect/cmd/protoc-gen-connect-go@latest",
	}
	for _, cmd := range commands {
		if err := sh.Run("go", "install", cmd); err != nil {
			return fmt.Errorf("failed to install command %s: %w", cmd, err)
		}
	}
	return nil
}

// Test runs the Go tests in the codebase.
func (Go) Test() error {
	return sh.Run("go", "test", "./...")
}

// Build compiles the Go codebase into a binary for windows/amd64.
// The binary will be placed in .bin/ directory.
func (Go) Build() error {
	const (
		goos = "windows"
		arch = "amd64"
		bin  = ".bin"
	)
	if err := os.MkdirAll(bin, 0o755); err != nil {
		return fmt.Errorf("failed to create %s directory: %w", bin, err)
	}
	commands := []string{
		"usi-bridge",
		"usi-mcp-server",
	}
	for _, cmd := range commands {
		executable := fmt.Sprintf("%s/%s-%s-%s.exe", bin, cmd, goos, arch)
		main := fmt.Sprintf("./cmd/%s", cmd)
		env := map[string]string{
			"GOOS":   goos,
			"GOARCH": arch,
		}
		if err := sh.RunWith(env, "go", "build", "-o", executable, main); err != nil {
			return fmt.Errorf("failed to build %s: %w", cmd, err)
		}
		fmt.Printf("Built %s to %s\n", cmd, executable)
	}
	return nil
}
