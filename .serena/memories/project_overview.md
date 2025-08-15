Project Purpose:
- shogithink is a modular system that combines Shogi AI with AI-driven commentary to explain positions, best moves, and evaluations in natural language.
- The current main component is usi-bridge, a simple USI engine for Shogi applications.

Tech Stack:
- Go (main language)
- Mage (task runner)
- Some Python support for environment setup (via Mage)

Codebase Structure:
- cmd/usi-bridge/: Main Go implementation for the USI engine
- magefiles/mage.go: Mage tasks for Go and Python
- go.mod/go.sum: Go dependencies

Entrypoints:
- usi-bridge: Run as a USI engine (stdin/stdout), also starts HTTP server on port 8080 for logs
- Mage tasks: Used for build, test, format, etc.

