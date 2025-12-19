# Contributing Guide

Thank you for contributing to the commitmate project! 🎉

## Development Environment Setup

### Prerequisites

- Go 1.21 or higher
- Git
- (Optional) GoReleaser (for release testing)

### Local Development

```bash
# Clone the repository
git clone https://github.com/leehosu/commitmate.git
cd commitmate

# Install dependencies
go mod download

# Build
go build -o commitmate

# Run
./commitmate --help
```

## Contributing Code

### Branch Strategy

- `main`: Stable version
- `feature/*`: New features
- `fix/*`: Bug fixes
- `docs/*`: Documentation changes

### Pull Request Process

1. Create an issue first to discuss the changes
2. Fork and create a branch
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. Commit your changes (using Conventional Commits format)
   ```bash
   git commit -m "feat: add amazing feature"
   ```
4. Push the branch
   ```bash
   git push origin feature/amazing-feature
   ```
5. Create a Pull Request

### Commit Message Rules

We follow the Conventional Commits format:

```
<type>(<scope>): <subject>

[optional body]

[optional footer]
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code formatting
- `refactor`: Refactoring
- `test`: Test additions/modifications
- `chore`: Build, configuration, etc.

**Examples:**
```
feat(ai): add support for GPT-4o-mini
fix(config): resolve config file parsing error
docs(readme): update installation instructions
```

## Testing

```bash
# Run unit tests
go test ./...

# Test specific package
go test ./internal/ai

# Check coverage
go test -cover ./...
```

## Code Style

- Format with `gofmt`
- Lint with `golangci-lint`
- Use clear and meaningful variable/function names
- Comments can be written in Korean or English

## Reporting Issues

Found a bug? Please create an issue!

**Include:**
- Expected behavior vs actual behavior
- Steps to reproduce
- Environment information (OS, Go version, etc.)
- Error messages and logs

## Adding a New AI Provider

To add a new AI provider:

1. Create a new file in `internal/ai/` directory (e.g., `gemini.go`)
2. Implement the `Client` interface
3. Add the new provider to `NewClient()` function
4. Add the new provider to the configuration struct
5. Write tests
6. Update README

## License

Your contributed code will be distributed under the MIT license.

## Have Questions?

Please create an issue or start a discussion!
