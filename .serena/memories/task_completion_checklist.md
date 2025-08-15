When a development task is completed:
- Run mage format to ensure code is properly formatted.
- Run mage go:tidy to clean up dependencies.
- Run mage go:test to verify all tests pass.
- Optionally run mage to perform all steps.
- Commit changes using git with a clear message.
- Push to the appropriate branch.

