[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/crazygreenpenguin/logger)
[![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/crazygreenpenguin/logger/main/LICENSE)
[![Build Status](https://travis-ci.com/crazygreenpenguin/logger.svg?branch=main)](https://travis-ci.com/crazygreenpenguin/logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/crazygreenpenguin/logger)](https://goreportcard.com/report/github.com/crazygreenpenguin/logger)
# logger
Simple logger for go

# Performance tests
```
goos: darwin
goarch: amd64
pkg: github.com/crazygreenpenguin/logger
BenchmarkInfof-4      	  650191	      1597 ns/op	     432 B/op	      11 allocs/op
BenchmarkInfo-4       	  736945	      1565 ns/op	     432 B/op	      11 allocs/op
BenchmarkErrorf-4     	  726738	      1629 ns/op	     432 B/op	      11 allocs/op
BenchmarkError-4      	  749658	      1589 ns/op	     432 B/op	      11 allocs/op
BenchmarkWarningf-4   	  703910	      1643 ns/op	     432 B/op	      11 allocs/op
BenchmarkWarning-4    	  753482	      1602 ns/op	     432 B/op	      11 allocs/op
```