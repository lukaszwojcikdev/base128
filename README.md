<div align="center">

```
  __________                          ____ ________    ______
  \______   \_____     ______  ____  /_   |\_____  \  /  __  \
   |    |  _/\__  \   /  ___/_/ __ \  |   | /  ____/  >      <
   |    |   \ / __ \_ \___ \ \  ___/  |   |/       \ /   --   \
   |______  /(____  //____  > \___  > |___|\_______ \\______  /
          \/      \/      \/      \/               \/       \/
```

# Base128

**A simple, open-source tool for encoding and decoding text data — available in Go and Python.**

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Website](https://img.shields.io/website?url=https%3A%2F%2Fwww.base128.pl)](http://www.base128.pl)
[![Version](https://img.shields.io/badge/version-2.1-blue)](#)
[![Go](https://img.shields.io/badge/-Go-00ADD8?logo=go&logoColor=white&style=flat)](https://go.dev)
[![Python](https://img.shields.io/badge/-Python%203.10%2B-3776AB?logo=python&logoColor=white&style=flat)](https://www.python.org)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](#contributing)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-blue?logo=linkedin)](https://www.linkedin.com/in/lukasz-michal-wojcik)

</div>

---

Base128 works by numerically representing input data with 7-bit groups
(values 0–127), enabling a simple, human-inspectable numeric representation
of text — useful for education, data representation experiments, and
lightweight text transformation tasks.

> ⚠️ Base128 is **not** a security, encryption, or compression tool. It only
> transforms text into a sequence of integers and back.

## Contents

- [Features](#features)
- [Application](#application)
- [Limitations](#limitations)
- [Installation](#installation)
  - [Go version](#go-version)
  - [Python version](#python-version)
- [Usage](#usage)
  - [Go](#go-usage)
  - [Python](#python-usage)
- [CLI options](#cli-options)
- [Examples](#examples)
- [Contributing](#contributing)
- [Documentation](#documentation)
- [Author](#author)
- [Site](#site)
- [Download](#download)
- [License](#license)

## Features

- ✅ Encodes and decodes plain text files (`*.txt`, `*.svg`, `*.html`, ...)
- ✅ Available in two independent, functionally identical implementations:
  **Go** and **Python 3**
- ✅ Reads from a file or from standard input, writes to standard output
- ✅ Clear error handling for out-of-range characters and malformed input
- ✅ Optional `--preserve-newlines` flag to keep original line breaks
- ✅ Zero external dependencies — both versions use only the standard library

## Application

Typical uses of Base128:

**Text data transformation** — turn text into a sequence of integers, making
it less directly human-readable. Example: encoding `"Hello"` produces
`72 101 108 108 111`.

**Data representation** — store or transmit text as a list of integers,
useful in specific data-processing or storage scenarios.

**Educational purposes** — a compact, readable example of how text
characters map to their ASCII values and back.

## Limitations

- **Text-only support** — designed for text data; binary data is not
  supported and may produce errors or incorrect output.
- **No compression or security** — the algorithm neither compresses data nor
  provides encryption. It only converts text to integers and back.
- **Limited character range** — works with characters in the range 0–127
  (ASCII); it does not handle extended character sets or binary data.

## Installation

### Go version

```bash
# 1. Clone the repository
git clone https://github.com/lukaszwojcikdev/base128.git
cd base128

# 2. Compile the source code
go build base128.go

# 3. Ready to use
./base128       # Linux / macOS
base128.exe     # Windows
```

### Python version

Requires **Python 3.10+**. No external dependencies.

```bash
# 1. Clone the repository (if not already done)
git clone https://github.com/lukaszwojcikdev/base128.git
cd base128

# 2. Run directly
python3 base128.py -h

# 3. (Optional) make it executable on Linux/macOS
chmod +x base128.py
./base128.py -h
```

## Usage

### Go usage

```bash
# Encode
./base128 -e data.txt > encoded.txt

# Decode
./base128 -d encoded.txt > decoded.txt
```

### Python usage

```bash
# Encode
python3 base128.py -e data.txt > encoded.txt

# Decode
python3 base128.py -d encoded.txt > decoded.txt

# Encode from standard input
echo "Hello, World!" | python3 base128.py -e

# Preserve newlines when encoding/decoding
python3 base128.py -e --preserve-newlines data.txt > encoded.txt
python3 base128.py -d encoded.txt > decoded.txt
```

Both implementations share the same command-line interface, so scripts and
examples below apply equally to `./base128` and `python3 base128.py`.

## CLI options

| Flag                    | Description                                                      |
| ------------------------ | ----------------------------------------------------------------- |
| `-e`, `--encode`         | Encodes the input into Base128 format.                            |
| `-d`, `--decode`         | Decodes data previously encoded with Base128.                     |
| `-h`, `--help`           | Prints usage instructions and the list of available options.      |
| `--preserve-newlines`    | Keeps `\n` / `\r` characters instead of stripping them.           |
| `--version`              | Prints the program's version.                                     |
| `--copyright`            | Prints copyright information.                                     |

## Examples

**Encoding a message:**

```text
$ echo -n "Hello, World!" | python3 base128.py -e
72 101 108 108 111 44 32 87 111 114 108 100 33
```

**Decoding it back:**

```text
$ echo "72 101 108 108 111 44 32 87 111 114 108 100 33" | python3 base128.py -d
Hello, World!
```

**Handling invalid input:**

```text
$ echo "72 200 108" | python3 base128.py -d
base128: liczba na pozycji 1 (200) jest poza zakresem Base128 (0-127)
```

## Contributing

Contributions are always welcome! To contribute to Base128:

1. Fork the repository on GitHub.
2. Create a new branch for your feature or bugfix.
3. Make your changes and ensure both the Go and Python implementations stay
   in sync in behavior (options, error messages, exit codes).
4. Submit a pull request with a clear description of your changes.

See [CONTRIBUTING.md](CONTRIBUTING.md) and
[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details.

## Documentation

- [Base128 source code (PDF)](http://www.base128.pl/base128_source_code.pdf)

## Author

- [@lukaszwojcikdev](https://www.github.com/lukaszwojcikdev)

## Site

- [www.base128.pl](http://www.base128.pl)
- [www.lukaszwojcik.eu](http://www.lukaszwojcik.eu)

## Download

| Platform     | Link                                                     |
| ------------ | --------------------------------------------------------- |
| Source (ZIP) | [Download](http://www.base128.pl/base128.zip)             |
| MD5 checksum | [Download](http://www.base128.pl/base128md5sum.md5)       |
| Go source    | [base128.go](base128.go)                                  |
| Python source| [base128.py](base128.py)                                  |

## License

[MIT License](LICENSE.md)

<div align="center">

Made with ❤️ by [Łukasz Wójcik](http://www.lukaszwojcik.eu) &middot; [www.base128.pl](http://www.base128.pl)

</div>
