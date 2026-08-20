# Log Aggregator

![CI](https://github.com/Qyroxen/Log-Aggregator/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Log-Aggregator/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Log-Aggregator?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Log-Aggregator)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Log-Aggregator)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Log-Aggregator?style=social)](https://github.com/Qyroxen/Log-Aggregator/stargazers)

## What is it?

Log Aggregator is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Log-Aggregator.git
cd Log-Aggregator
go build -o logaggregator .

# Run
./logaggregator --help
```

## CLI Usage

```bash
# Basic usage
./logaggregator

# With flags
./logaggregator --verbose --output json

# Get help
./logaggregator --help
```

## Examples

```bash
# Example 1
./logaggregator example1

# Example 2
./logaggregator example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o logaggregator .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Log-Aggregator/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Log-Aggregator?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Log-Aggregator/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Log-Aggregator?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Log-Aggregator/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Log-Aggregator" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Log-Aggregator/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Log-Aggregator" alt="Pull Requests">
  </a>
</p>
