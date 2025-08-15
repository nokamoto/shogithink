# Project Purpose
A modular system combining Shogi AI with AI-driven commentary to explain positions, best moves, and evaluations in natural language.

# Tech Stack
- Go (main language)
- Mage (task automation)
- Python (environment and package management)
- `uv` package
- MCP server integration via `.vscode/mcp.json`
- Linux environment

# Code Structure
- Go source files
- `magefiles/mage.go` for Mage automation tasks
- Python virtual environment in `.venv`
- `.vscode/mcp.json` for MCP server configuration

# License
MIT License

# Mage Commands
- `mage python:venv`: Create Python virtual environment
- `mage python:install`: Install `uv` package

# Environment
- Linux
- Uses `.venv` for Python

# Missing Information
- Code style/conventions (Go and Python)
- Guidelines, design patterns, or conventions
- Commands for linting, formatting, and testing
- Entrypoints for running the project
- What to do when a task is completed

Further investigation needed for conventions and commands.