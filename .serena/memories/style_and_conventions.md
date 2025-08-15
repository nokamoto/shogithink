
# Style and Conventions

## Code Style

- Go code follows standard Go conventions.
- Formatting is enforced using goimports and gofumpt (see Format task in `magefiles/mage.go`).
- No explicit mention of type hints or docstrings beyond standard Go comments.
- Python code (if any) uses standard Python conventions, with venv setup via Mage.

## Design Patterns

- Simple command-line and HTTP server pattern for usi-bridge.
- Mage is used for task automation.
