# commitmate

🤖 AI-powered Git commit message generator

**[English](README.md)** | [한국어](docs/ko.md)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/leehosu/commitmate)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/leehosu/commitmate)](https://github.com/leehosu/commitmate/releases)

>  **[Contributing](CONTRIBUTING.md)** | **[Changelog](CHANGELOG.md)**

## Features

- ✨ **AI-powered commit message generation**: Supports OpenAI GPT and Anthropic Claude
- 📝 **Conventional Commits format**: Industry-standard commit message convention
- 🌏 **Multilingual support**: Korean/English commit messages and UI
- 🎫 **JIRA integration**: Auto-add JIRA issue numbers from branch names
- 🎯 **Simple usage**: Commit with a single command
- ⚙️ **Flexible configuration**: Choose API keys and providers
- 🚀 **Cross-platform**: Linux, macOS, Windows support

## Installation

### Homebrew (Recommended) 🍺

```bash
# Add tap
brew tap leehosu/tap

# Install
brew install commitmate

# Verify
commitmate version
```

### Binary Download

Download the binary for your OS from the latest release:
[Releases](https://github.com/leehosu/commitmate/releases)

```bash
# macOS/Linux
tar -xzf commitmate_*_*.tar.gz
chmod +x commitmate
sudo mv commitmate /usr/local/bin/

# Windows
# Extract commitmate.exe and add to PATH
```

## Quick Start

### 1. Set API Key

**Using OpenAI:**
```bash
commitmate config set-key openai sk-xxxxx
commitmate config set-provider openai
```

**Using Claude:**
```bash
commitmate config set-key claude sk-ant-xxxxx
commitmate config set-provider claude
```

### 2. Generate Commit

```bash
# After making changes
git add .

# AI automatically generates commit message and commits
commitmate
```

## Usage

### Basic Commands

```bash
# Basic usage (analyze staged changes and commit)
commitmate

# Generate message only without committing
commitmate --dry-run

# Use specific AI provider (one-time)
commitmate --provider openai
commitmate --provider claude

# Skip git hooks
commitmate --no-verify
```

### Configuration Management

```bash
# Set API key
commitmate config set-key openai sk-xxxxx
commitmate config set-key claude sk-ant-xxxxx

# Set default provider
commitmate config set-provider openai

# Change model
commitmate config set-model openai gpt-4o-mini
commitmate config set-model claude claude-3-5-haiku-20241022

# Language settings
commitmate config set-commit-language ko  # Commit message language (ko/en)
commitmate config set-ui-language en      # UI language (ko/en)

# Show current configuration
commitmate config show

# Check version
commitmate version
```

### Environment Variables

You can also configure using environment variables:

```bash
export COMMITMATE_OPENAI_API_KEY=sk-xxxxx
export COMMITMATE_CLAUDE_API_KEY=sk-ant-xxxxx
export COMMITMATE_PROVIDER=openai
export COMMITMATE_COMMIT_LANGUAGE=ko  # Commit message language
export COMMITMATE_UI_LANGUAGE=en      # UI language
```

## Conventional Commits

commitmate follows the [Conventional Commits](https://www.conventionalcommits.org/) format:

```
<type>(<scope>): <subject>

[optional body]

[optional footer]
```

**Supported types:**
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Code formatting (no functional changes)
- `refactor`: Code refactoring
- `test`: Adding/modifying tests
- `chore`: Build, config, and other changes
- `perf`: Performance improvements
- `ci`: CI configuration changes
- `build`: Build system changes
- `revert`: Revert previous commit

## Multilingual Support

commitmate supports both Korean and English for commit messages and UI:

```bash
# Commit message language (default: en)
commitmate config set-commit-language en  # or ko

# UI language (default: ko)
commitmate config set-ui-language ko      # or en
```

## JIRA Integration

commitmate **automatically** detects JIRA issue numbers from branch names and adds them to commit messages - no configuration needed!

**Example:**
```bash
# Branch with JIRA issue
git checkout -b DEVOPS2-430-add-user-feature
commitmate
# Result: [DEVOPS2-430] feat: add user authentication

# Branch without JIRA issue
git checkout -b feature/add-auth
commitmate
# Result: feat: add user authentication
```

**Supported patterns:** `PROJECT-123`, `ABC-456`, `DEVOPS2-430`

**Note:** JIRA prefixes are not added on main, master, or develop branches.

## Example

```bash
$ git add .
$ commitmate

🔍 Analyzing staged changes...
✨ AI generated commit message:

feat(auth): add JWT authentication middleware

? Do you want to use this commit message? 
  ▸ Yes - commit
    Edit - edit and commit
    Regenerate - generate again
    Cancel

✓ Commit completed successfully!
```


## License

MIT License - See [LICENSE](LICENSE) file

## Contributing

Issues and PRs are welcome!

## Author

[@leehosu](https://github.com/leehosu)
