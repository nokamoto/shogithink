package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = All

// Python provides Mage targets for Python environment setup and package installation.
type Python mg.Namespace

// Venv creates a Python virtual environment in the .venv directory.
func (Python) Venv() error {
	return sh.Run("python3", "-m", "venv", ".venv")
}

// Install installs the 'uv' package using pip.
func (Python) Install() error {
	return sh.Run("pip", "install", "uv")
}

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

// Format runs formatters on the codebase.
func Format() error {
	formatters := [][]string{
		{"goimports", "-l", "-w", "."},
		{"gofumpt", "-l", "-w", "."},
		{"yamlfmt", "."},
	}
	for _, formatter := range formatters {
		if err := sh.Run(formatter[0], formatter[1:]...); err != nil {
			return fmt.Errorf("failed to run formatter %s: %w", formatter[0], err)
		}
	}
	return nil
}

// All is the default target that runs all Mage tasks.
func All() {
	mg.SerialDeps(
		Go.Install,
		Format,
		Go.Tidy,
		Go.Test,
		Go.Build,
	)
}
