# Contributing to SupaDash

Thank you for your interest in contributing! Here's how to get started.

## Development Setup

1. **Prerequisites**: Go 1.25+, Docker, PostgreSQL
2. Clone the repo and install dependencies:
   ```bash
   git clone https://github.com/tanviralamtusar/SupaDash.git
   cd SupaDash
   cp .env.example .env
   go mod download
   ```
3. Run locally:
   ```bash
   go run main.go
   ```

## Making Changes

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes
4. Run tests: `go test ./... -v`
5. Ensure the build passes: `go build ./...`
6. Commit with a descriptive message
7. Push and open a Pull Request

## Code Standards

- **Formatting**: Run `gofmt` on all Go files
- **Naming**: Follow Go naming conventions
- **Tests**: Add tests for new functionality
- **Comments**: Document exported functions and types

## Pull Request Process

1. Update documentation if your change affects the API or configuration
2. Add or update tests as needed
3. Ensure CI passes (tests + build)
4. Request review from a maintainer

## Reporting Bugs

Open a [GitHub Issue](https://github.com/tanviralamtusar/SupaDash/issues) with:
- Steps to reproduce
- Expected vs. actual behavior
- Go version, OS, and Docker version

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
