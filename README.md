# log [![Build Status](https://travis-ci.org/vinxi/log.png)](https://travis-ci.org/vinxi/log) [![GoDoc](https://godoc.org/github.com/vinxi/log?status.svg)](https://godoc.org/github.com/vinxi/log) [![Coverage Status](https://coveralls.io/repos/github/vinxi/log/badge.svg?branch=master)](https://coveralls.io/github/vinxi/log?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/vinxi/log)](https://goreportcard.com/report/github.com/vinxi/log)

Structured, versatile, hierarchical logging layer for vinxi's components. 

Built on top of [Logrus](https://github.com/Sirupsen/logrus).

## Installation

```bash
go get -u gopkg.in/vinxi/log.v0
```

## API

See [godoc](https://godoc.org/github.com/vinxi/log) reference.

## Example

#### Basic logging in middleware handlers

```go
package main

import (
  "fmt"
  "net/http"

  "gopkg.in/vinxi/log.v0"
  "gopkg.in/vinxi/vinxi.v0"
)

const port = 3100

func main() {
  // Create a new vinxi proxy
  vs := vinxi.NewServer(vinxi.ServerOptions{Port: port})

  // Plugin multiple middlewares writting some logs
  vs.Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
    log.Infof("[%s] %s", r.Method, r.RequestURI)
    h.ServeHTTP(w, r)
  })

  vs.Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
    log.Warnf("%s", "foo bar")
    h.ServeHTTP(w, r)
  })

  // Target server to forward
  vs.Forward("http://httpbin.org")

  fmt.Printf("Server listening on port: %d\n", port)
  err := vs.Listen()
  if err != nil {
    fmt.Errorf("Error: %s\n", err)
  }
}

```

## License

MIT
