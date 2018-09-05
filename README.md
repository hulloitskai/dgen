# dgen

_A command-line tool designed for repeating a string an excessive number of
times._

[![godoc: reference][godoc-img]][godoc]
[![go report card][goreportcard-img]][goreportcard]
[![travis: build][travis-img]][travis]
[![appveyor: build][appveyor-img]][appveyor]
[![codecov: coverage][codecov-img]][codecov]

## Installation

```bash
go get github.com/steven-xie/dgen
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

| Name / ID | Value  | service            |
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
import (
  "github.com/steven-xie/dgen/throughput"
  "os"
)


func main() {
  const (
    repstr  = "test string "
    reps    = 5000
    bufsize = throughput.RecommendedBufSize
  )

  // Dump "test string " 5000 times into os.Stdout.
  throughput.Dump(repstr, reps, bufsize, os.Stdout)
}
```

[godoc-img]: https://godoc.org/github.com/steven-xie/dgen?status.svg
[godoc]: https://godoc.org/github.com/steven-xie/dgen
[travis-img]: https://travis-ci.org/steven-xie/dgen.svg?branch=master
[travis]: https://travis-ci.org/steven-xie/dgen
[codecov-img]: https://codecov.io/gh/steven-xie/dgen/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/steven-xie/dgen
[appveyor]: https://ci.appveyor.com/project/StevenXie/dgen-p2xgh
[appveyor-img]: https://ci.appveyor.com/api/projects/status/wrry9krp1kkx4hsi/branch/master?svg=true
[goreportcard]: https://goreportcard.com/report/github.com/steven-xie/dgen
[goreportcard-img]: https://goreportcard.com/badge/github.com/steven-xie/dgen
