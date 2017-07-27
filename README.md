Hiro
====

Generates HTML from API Blueprints using the [Snow Crash](https://github.com/apiaryio/snowcrash) command line tool
[Drafter](https://github.com/apiaryio/drafter) and [Iglo](https://github.com/subosito/iglo).

[![GoDoc](https://godoc.org/github.com/peterhellberg/hiro/hiro?status.svg)](https://godoc.org/github.com/peterhellberg/hiro/hiro)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/hiro#license-mit)

## Dependencies

Hiro requires [Drafter](https://github.com/apiaryio/drafter) v3.2.2 to be installed.
Any newer release is not supported due to usage of the deprecated `--type=ast`
option ([apiaryio/drafter@a7dd4c5](https://github.com/apiaryio/drafter/commit/a7dd4c5)).

Install it on OS X using Homebrew:

```bash
$ brew install --HEAD \
  https://raw.githubusercontent.com/apiaryio/drafter/stable/3.x/tools/homebrew/drafter.rb
```

Refer to the Drafter [build](https://github.com/apiaryio/drafter#build) notes if you are using a different OS.

## Installation

```bash
$ go get -u github.com/peterhellberg/hiro
```

You can also download binaries from the [releases page](https://github.com/peterhellberg/hiro/releases).

## Usage

```bash
$ hiro --help

Command line arguments:

  -input="blueprint.md": Input file (.md)
  -output="index.html": Output file (.html)
  -template="": Iglo template file
```

## Example

```markdown
# Service Name

A short description of the service.

## Group Group 1

A description of this group of endpoints

## /endpoint

+ Parameters
  + foo (required, string, `example`)
  A description of the foo parameter
  + bar (optional, integer, `123`)
  A description of the bar parameter

### GET

A description of the endpoint.

header | another header
-------|---------------
foo    | bar
baz    | qux

> A quote related to the endpoint

**Bold text**

    func some(code Example) string {
      return "related to the endpoint"
    }

+ Response 200 (application/json; charset=utf-8)
  + Headers
  X-Example header
  + Body
  {"endpoint": "response"}

## Group Group 2

A description of this group of endpoints.
```

![Hiro Blueprint Output](http://assets.c7.se/skitch/Hiro_Blueprint_Output_Example-20160511-161557.png)

## License (MIT)

Copyright (c) 2015-2016 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
