package main

import (
	"fmt"

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
		Buf.Format,
		Buf.Gen,
		Format,
		Go.Tidy,
		Go.Test,
		Go.Build,
	)
}
