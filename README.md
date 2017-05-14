# [nebulo-golib](https://github.com/krostar/nebulo-golib) [![License](https://img.shields.io/github/license/krostar/nebulo-golib.svg)](https://tldrlegal.com/license/gnu-general-public-license-v3-(gpl-3)) [![Godoc](https://godoc.org/github.com/krostar/nebulo-golib?status.svg)](https://godoc.org/github.com/krostar/nebulo-golib)

This is the library used for Nebulo client and server.
/!\\ This project is a school project, it's not finished, it's not 100% working, don't use it unless you know what you do /!\\

[![Build status](https://travis-ci.org/krostar/nebulo-golib.svg?branch=dev)](https://travis-ci.org/krostar/nebulo-golib) [![Go Report Card](https://goreportcard.com/badge/github.com/krostar/nebulo)](https://goreportcard.com/report/github.com/krostar/nebulo-golib) [![Codebeat status](https://codebeat.co/badges/54741d30-dff6-45e1-bee4-13004944d118)](https://codebeat.co/projects/github-com-krostar-nebulo-golib-dev) [![Coverage status](https://coveralls.io/repos/github/krostar/nebulo-golib/badge.svg?branch=dev)](https://coveralls.io/github/krostar/nebulo-golib?branch=dev)

## Documentation
The Golang documentation is available on the [godoc website](https://godoc.org/github.com/krostar/nebulo-golib)

## Licence
Distributed under GPL-3 License, please see license file, and/or browse [tldrlegal.com](https://tldrlegal.com/license/gnu-general-public-license-v3-(gpl-3)) for more details.

## Contribute to the project
### Report bugs
Create an [issue](https://github.com/krostar/nebulo-golib/issues) or contact [bug[at]nebulo[dot]io](mailto:bug@nebulo.io)

### Before you started
#### Check your golang installation
Make sure `golang` is installed and is at least in version **1.8** and your `$GOPATH` environment variable set in your working directory
```sh
$> go version
go version go1.8 linux/amd64
$> echo $GOPATH
/home/krostar/go
```

If you don't have `golang` installed or if your `$GOPATH` environment variable isn't set, please visit [Golang: Getting Started](https://golang.org/doc/install) and [Golang: GOPATH](https://golang.org/doc/code.html#GOPATH)

> It may be a good idea to add `$GOPATH/bin` and `$GOROOT/bin` in your `$PATH` environment!

#### Download the project
```sh
# Manually
$> mkdir -p $GOPATH/src/github.com/krostar/
$> git -c $GOPATH/src/github.com/krostar/ clone https://github.com/krostar/nebulo-golib.git

# or via go get
$> go get github.com/krostar/nebulo-golib
```

#### Download the tool manager
```sh
$> go get -u github.com/twitchtv/retool
```

#### Use our Makefile
We are using a Makefile to everything we need (build, release, tests, documentation, ...).
```sh
# Get the dependencies and tools
$> make vendor

# Build the project (by default generated binary will be in <root>/build/bin/nebulo)
$> make build

# Test the project
$> make test
```

### Guidelines
#### Coding standart
Please, make sure your favorite editor is configured for this project. The source code should be:
- well formatted (`gofmt` (usage of tabulation, no trailing whitespaces, trailing line at the end of the file, ...))
- linter free (`gometalinter --config=.gometalinter.json ./...`)
- with inline comments beginning with a lowercase caracter

Make sure to use `make test` before submitting a pull request!

### Other things
- use the dependencies manager and update them (see [govendor](https://github.com/kardianos/govendor) and [retool](https://github.com/twitchtv/retool))
- write unit tests
