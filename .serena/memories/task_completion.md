# Updated Task Completion Workflow

1. Ensure code is formatted (`mage format` for Go, PEP8 for Python).
2. Run linting (`go vet` for Go, flake8 for Python if used).
3. Run tests (`go test` for Go, pytest for Python if used).
4. Update documentation if needed.
5. Commit changes using `git`.
6. Push to remote repository.
7. If MCP server is involved, ensure `.vscode/mcp.json` is updated and server is restarted if necessary.
8. Use `mage all` to run all setup and formatting tasks.

Update as more workflow steps are discovered.