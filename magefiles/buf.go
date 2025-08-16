package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Buf provides Mage targets for Buf-related tasks.
type Buf mg.Namespace

// Format formats the protobuf files in the codebase.
func (Buf) Format() error {
	return sh.Run("buf", "format", "-w", ".")
}

// Gen generates code from the protobuf files.
func (Buf) Gen() error {
	return sh.Run("buf", "generate")
}
