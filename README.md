[![CI](https://github.com/userosettadev/rosetta-cli/actions/workflows/go.yml/badge.svg)](https://github.com/userosettadev/rosetta-cli/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/userosettadev/rosetta-cli)](https://goreportcard.com/report/github.com/userosettadev/rosetta-cli)
[![codecov](https://codecov.io/gh/userosettadev/rosetta-cli/graph/badge.svg?token=XOWMEPON83)](https://codecov.io/gh/userosettadev/rosetta-cli)

# Getting Started
Rosetta is a service that converts code into OpenAPI specification.

### Limitations
- **Supported Languages**: Currently, _Rosetta_ exclusively supports Go. However, we are actively working on expanding language support. Stay tuned for updates :)
- **Number of Runs**: The number of runs allowed is based on your account type.
- **Project Size**: The maximum project size is based on your account type.

## Installation

### Install with Go
```bash
go install github.com/userosettadev/rosetta-cli@latest
```

### macOS
Rosetta can be installed on macOS using [Homebrew](https://brew.sh/):
```bash
brew tap userosettadev/rosetta
brew install rosetta
```

### Windows
For Windows, you can download the Rosetta executable from the [releases page](TODO).

1. Visit the releases page and download the latest version of Rosetta for Windows.
2. Extract the downloaded archive.
3. Move the `rosetta.exe` file to a directory in your system's `PATH`.

## Verify Installation
After the installation is complete, you can verify it by running the following command (or executing the appropriate method for your installation):
```
rosetta --version
```

## Setup
Before using Rosetta, you need to set the `ROSETTA_API_KEY` environment variable:
```
export ROSETTA_API_KEY=<your_api_key>
```

## Running Rosetta - Generating an OpenAPI Spec
To generate an OpenAPI specification from your code using Rosetta, you can use the `gen` command followed by the path to your code and the programming language.
```
rosetta gen /path/to/code -l <language>
```

### Example
To generate an OpenAPI specification from Go code located in the `./myapp` directory, you would run:
```
rosetta gen ./myapp -l go
```
This command will analyze the Go code in the `./myapp` directory and generate an OpenAPI specification based on the code.

### Docker
Rosetta is also available as a Docker image:
```bash
docker run --rm -v $PWD:/app -w /app -e ROSETTA_API_KEY=$ROSETTA_API_KEY ghcr.io/userosettadev/rosetta-cli gen /path/to/code -l go
