# dgen

_A command-line tool designed for repeating a string an excessive number of
times._

[![godoc: reference][godoc-img]][godoc]
[![travis: build][travis-img]][travis]
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

[godoc-img]: https://godoc.org/github.com/steven-xie/dgen?status.svg
[godoc]: https://godoc.org/github.com/steven-xie/dgen
[travis-img]: https://travis-ci.org/steven-xie/dgen.svg?branch=master
[travis]: https://travis-ci.org/steven-xie/dgen
[codecov-img]: https://codecov.io/gh/steven-xie/dgen/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/steven-xie/dgen
