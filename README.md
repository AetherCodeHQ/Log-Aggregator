# 🌐 Log Aggregator

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Infrastructure tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`infrastructure` `devops` `cli` `golang` `io`

---

## What is Log-Aggregator?

**Log-Aggregator** is an infrastructure tool for monitoring, inspecting, and managing systems and services.

## Features

- ✅ Streaming file processing
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Log-Aggregator.git
cd Log-Aggregator

# Build
go build -o log-aggregator .

# Run
./log-aggregator Usage: log-aggregator <logfile>
```

### Or directly with `go run`:
```bash
go run main.go Usage: log-aggregator <logfile>
```

## Usage

```bash
# Basic usage
./log-aggregator Usage: log-aggregator <logfile>

# With flags
./log-aggregator Usage: log-aggregator <logfile> value Usage: log-aggregator <logfile>
```

### Example Output

```
$ ./log-aggregator Usage: log-aggregator <logfile>
Usage: log-aggregator <logfile>
Error:
Log Aggregator Summary
```

## Project Structure

```
Log-Aggregator/
  main.go          # Entry point (57 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
