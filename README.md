# go-readline

go-readline is a wrapper for the
[GNU Readline library](http://cnswww.cns.cwru.edu/php/chet/readline/rltop.html)
for the [Go programming language](http://golang.org).

This repository contains work from multiple contributors. See the AUTHORS file
for details.


## Requirements

go-readline requires Go 1.2 or newer. Download the latest from:
https://code.google.com/p/go/downloads/

Git is required for installing via `go get`.

A C compiler (gcc or clang) and the GNU Readline library are required too.

On Debian and Ubuntu, install the development packages:

	apt-get install build-essential libreadline-dev

On CentOS and RHEL:

	yum install gcc readline-devel

On Mac OS X, via [homebrew](http://brew.sh):

	brew install readline

## Installation

Make sure $GOROOT and [$GOPATH](http://golang.org/doc/code.html#GOPATH) are set, and install:

	go get github.com/fiorix/go-readline
