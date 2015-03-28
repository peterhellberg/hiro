Hiro
====

Generates HTML from API Blueprints using [Snow Crash](https://github.com/apiaryio/snowcrash) and [Iglo](https://github.com/subosito/iglo).

[![GoDoc](https://godoc.org/github.com/peterhellberg/hiro/hiro?status.svg)](https://godoc.org/github.com/peterhellberg/hiro/hiro)

## Installation

```bash
$ go get -u github.com/peterhellberg/hiro
```

## Dependencies

Hiro requires [Snow Crash](https://github.com/apiaryio/snowcrash) to be installed.

Install it on OS X using Homebrew:

```bash
$ brew install --HEAD \
  https://raw.github.com/apiaryio/snowcrash/master/tools/homebrew/snowcrash.rb
```

Refer to the snowcrash page if you are using a different OS.

## Usage

```bash
$ hiro --help

Command line arguments:

  -input="blueprint.md": Input file (.md)
  -output="index.html": Output file (.html)
  -template="": Iglo template file
```

## License

MIT License. See the [LICENSE](https://github.com/peterhellberg/hiro/blob/master/LICENSE) file.
