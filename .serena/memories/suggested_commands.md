Suggested Commands for Development:

# Go
- mage go:install      # Install Go tools (goimports, gofumpt, yamlfmt)
- mage format          # Format codebase
- mage go:tidy         # Clean up go.mod
- mage go:test         # Run Go tests
- mage go:build        # Build Windows binary for usi-bridge
- mage                 # Run all tasks (default: install, format, tidy, test, build)

# Python (if needed)
- mage python:venv     # Create Python virtual environment
- mage python:install  # Install 'uv' package

# Running usi-bridge
- go run ./cmd/usi-bridge/main.go   # Run USI engine (stdin/stdout, HTTP logs)

# Utilities
- git, ls, cd, grep, find           # Standard Linux utilities

