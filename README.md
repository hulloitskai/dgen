# dgen

_A CLI tool for repeating a string an excessive number of times._

[![Github Release][release-img]][release]
[![Go Report Card][grc-img]][grc]
[![Travis: Build][travis-img]][travis]
[![Codecov: Coverage][codecov-img]][codecov]

> Also check out [`begone`](https://github.com/stevenxie/begone), a fully
> automatic spamming tool for Facebook Messenger.

## Installation

If you're on macOS, you can get `dgen` on Homebrew:

```bash
brew install stevenxie/tap/dgen
```

Otherwise, check out the
[latest compiled release](https://github.com/stevenxie/dgen/releases)!

## Making from source

To install from source, make sure you have [Go](https://golang.org) installed,
and run:

```bash
## Clone the repository.
$ git clone git@github.com:stevenxie/dgen
$ cd dgen

## Install dependencies, and then install to $GOBIN.
$ make dl       # (or go mod download)
$ make install  # (or go install)
```

## Usage

To repeat "test string" 50 times, simply perform the following command:

```bash
dgen "test string " 50
# Output: test string test string (...)
```

For more advanced usage options, view the help prompt with:

```bash
dgen --help
```

### Presets

`dgen` comes with several presets in order to make your life easier. Each preset
represents the character limit on the associated messaging service:

```bash
dgen "👀 " fb
# Output: 👀 👀 👀 👀 👀 👀 👀 👀 👀 (... x5000)
```

#### List of current `dgen` presets:

| Name / ID | Value  | Service            |
| --------- | ------ | ------------------ |
| fb        | 5000   | Facebook Messenger |
| twitter   | 280    | Twitter            |
| rpost     | 40,000 | Reddit (post)      |
| rcomment  | 10,000 | Reddit (comment)   |
| rmsg      | 10,000 | Reddit (message)   |

## Performance

`dgen` is one of the top-of-the-line string generators out there. `dgen v1.0.1`
sports the following benchmark:

- **Input:** "benchmark test text "
- **Repetitions:** 1,000,000
- **Duration:** 25.050 seconds
- **Milliseconds / operation:** 0.025039
- **Bytes allocated / operation:** 21
- **Allocations / operation:** 0

## Integration

`dgen` is a thin wrapper around an internal library, `throughput`, which
contains the actual repeating and buffering logic. To create a program with
`dgen`'s string dumping capabilities, simply import `throughput`:

```go
package main

import (
  "github.com/stevenxie/dgen/throughput"
  "os"
)


func main() {
  var (
    repstr  = "test string "
    reps    = 5000
    bufsize = throughput.RecommendedBufSize
  )

  // Dump "test string " 5000 times into os.Stdout.
  throughput.Dump(repstr, reps, bufsize, os.Stdout)
}
```

<br />

# Development

If you want to help develop `dgen`, clone the repository and run `make setup`
to configure your development environment.

[release]: https://github.com/stevenxie/dgen/releases
[release-img]: https://img.shields.io/github/release/stevenxie/dgen.svg
[travis]: https://travis-ci.org/stevenxie/dgen
[travis-img]: https://travis-ci.org/stevenxie/dgen.svg?branch=master
[codecov]: https://codecov.io/gh/stevenxie/dgen
[codecov-img]: https://codecov.io/gh/stevenxie/dgen/branch/master/graph/badge.svg
[grc]: https://goreportcard.com/report/github.com/stevenxie/dgen
[grc-img]: https://goreportcard.com/badge/github.com/stevenxie/dgen
