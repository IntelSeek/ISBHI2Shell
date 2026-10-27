# ISBHI2Shell

**An independent command shell developed by ISBHI (IntelSeek Bureau of Hacking and Intelligence).**

ISBHI2Shell is a command shell written in Go, designed to provide its own built-in commands and a consistent command-line experience.

> **Status:** Early development — v0.1.0

## Features

- Custom command-line interface
- Built-in command system
- Dynamic prompt displaying the current user and directory
- Directory navigation
- File and directory listing
- Command-specific help
- Written in Go

## Built-in Commands

| Command | Description |
|---|---|
| `help` | Lists available commands |
| `help <command...>` | Displays help for specific commands |
| `version` | Displays the current version |
| `about` | Displays information about ISBHI2Shell |
| `whoami` | Displays the current username |
| `pwd` | Displays the current working directory |
| `cd [path]` | Changes the current working directory |
| `ls [path]` | Lists files and directories |
| `clear` | Clears the terminal |
| `exit` | Exits ISBHI2Shell |

## Requirements

- [Go](https://go.dev/dl/) installed on your system to run from source.

## Getting Started

### Run from source

Clone the repository:

```bash
git clone https://github.com/IntelSeek/ISBHI2Shell.git
cd ISBHI2Shell
```

Run the application:

```bash
go run .
```

### Build the executable

To build ISBHI2Shell for your current operating system:

```bash
go build -o ISBHI2Shell .
```

On Windows, this produces `ISBHI2Shell.exe`.

## Usage

Start ISBHI2Shell and type:

```text
help
```

To get information about a specific command:

```text
help version
help cd
help ls
```

## Project Goals

ISBHI2Shell aims to evolve into an independent, cross-platform command shell with its own built-in commands and consistent behavior across supported operating systems.

Cross-platform compatibility and functionality may vary during early development.

## License

ISBHI2Shell is distributed under the [MIT License](LICENSE).

## Developer

**ISBHI — IntelSeek Bureau of Hacking and Intelligence**
