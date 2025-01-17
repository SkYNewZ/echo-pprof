# echo-pprof
A wrapper for [Echo web framework](https://github.com/labstack/echo) to use `net/http/pprof` easily.

Forked from [sevenNt/echo-pprof](https://github.com/sevenNt/echo-pprof).
Support for Echo [v4](https://github.com/SkYNewZ/echo-pprof/releases/tag/v4) and [v4](https://github.com/SkYNewZ/echo-pprof/releases/tag/v5).

## Install

```sh
go get -u github.com/SkYNewZ/echo-pprof/v5
```

## Usage

```go
package main

import (
	"github.com/SkYNewZ/echo-pprof/v5"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	echopprof.Wrap(e)

	// echopprof also plays well with *echo.Group
	// prefix := "/debug/pprof"
	// group := e.Group(prefix)
	// echopprof.WrapGroup(prefix, group)

	e.Start(":8080")
}
```

Start this server, and then visit [http://127.0.0.1:8080/debug/pprof/](http://127.0.0.1:8080/debug/pprof/), and you'll
see the pprof index page.

Have fun.
