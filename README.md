<div align="center">

# 🤖 commitmate

**AI-powered Git commit message generator**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/leehosu/commitmate)](https://golang.org/)
[![Release](https://img.shields.io/github/v/release/leehosu/commitmate)](https://github.com/leehosu/commitmate/releases)

[English](README.md) | [한국어](docs/ko.md)

[Features](#features) • [Quick Start](#quick-start) • [Installation](#installation) • [Configuration](#configuration) • [Advanced](#advanced-features)

</div>

---

## ✨ Features

- 🤖 **AI-Powered** - OpenAI GPT & Anthropic Claude support
- 📝 **Conventional Commits** - Industry-standard commit format
- 🌏 **Multilingual** - Korean/English support for messages and UI
- 🎫 **JIRA Integration** - Auto-detect and add issue numbers from branch names
- 🎨 **Interactive UI** - Edit, regenerate, or cancel with simple prompts
- ⚙️ **Flexible** - Configurable via CLI or environment variables
- 🚀 **Cross-platform** - Linux, macOS, Windows

## 🚀 Quick Start

```bash
# 1. Set your API key
commitmate config set-key openai sk-xxxxx
commitmate config set-provider openai

# 2. Stage your changes
git add .

# 3. Generate and commit
commitmate
```

## 📖 Example

```bash
$ git add .
$ commitmate

🔍 Analyzing staged changes...
✨ AI generated commit message:

feat(auth): add JWT authentication middleware

? Do you want to use this commit message? 
  ▸ Yes - commit
    Edit - edit message
    Regenerate - generate again
    Cancel

✓ Commit completed successfully!
```

## 📦 Installation

### Homebrew (Recommended)

```bash
brew tap leehosu/tap
brew install commitmate
```

### Binary Download

Download from [Releases](https://github.com/leehosu/commitmate/releases)

```bash
# macOS/Linux
tar -xzf commitmate_*.tar.gz
chmod +x commitmate
sudo mv commitmate /usr/local/bin/

# Windows
# Extract commitmate.exe and add to PATH
```

## ⚙️ Configuration

### Basic Setup

```bash
# API Keys
commitmate config set-key openai sk-xxxxx
commitmate config set-key claude sk-ant-xxxxx

# Provider
commitmate config set-provider openai  # or claude

# Model (optional)
commitmate config set-model openai gpt-4o-mini
commitmate config set-model claude claude-3-5-haiku-20241022
```

### Language Settings

```bash
commitmate config set-commit-language en  # Commit message language (en/ko)
commitmate config set-ui-language ko      # UI language (en/ko)
```

### View Configuration

```bash
commitmate config show
```

### Environment Variables

```bash
export COMMITMATE_OPENAI_API_KEY=sk-xxxxx
export COMMITMATE_CLAUDE_API_KEY=sk-ant-xxxxx
export COMMITMATE_PROVIDER=openai
export COMMITMATE_COMMIT_LANGUAGE=ko
export COMMITMATE_UI_LANGUAGE=en
```

## 💻 Usage

```bash
commitmate                  # Analyze and commit
commitmate --dry-run        # Generate message only
commitmate --provider openai  # Use specific provider
commitmate --no-verify      # Skip git hooks
```

## 🎯 Advanced Features

### Conventional Commits

commitmate follows the [Conventional Commits](https://www.conventionalcommits.org/) format:

```
<type>(<scope>): <subject>
```

**Supported types:** `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`, `perf`, `ci`, `build`, `revert`

### JIRA Integration

Automatically detects JIRA issue numbers from branch names:

```bash
# Branch: DEVOPS2-430-add-feature
commitmate
# Output: [DEVOPS2-430] feat: add user authentication

# Branch: feature/add-auth
commitmate
# Output: feat: add user authentication
```

**Supported patterns:** `PROJECT-123`, `ABC-456`, `DEVOPS2-430`

**Note:** JIRA prefixes are not added on `main`, `master`, or `develop` branches.

### Multilingual Support

Separate language settings for commit messages and UI:

```bash
# English commits, Korean UI (for Korean developers in global teams)
commitmate config set-commit-language en
commitmate config set-ui-language ko

# Korean commits, English UI (for international developers in Korean companies)
commitmate config set-commit-language ko
commitmate config set-ui-language en
```

## 🤝 Contributing

Issues and PRs are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md)

## 📝 Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history

## 📄 License

MIT License - See [LICENSE](LICENSE)

## 👤 Author

[@leehosu](https://github.com/leehosu)

---

<div align="center">

**⭐ Star this project if you find it helpful!**

</div>
