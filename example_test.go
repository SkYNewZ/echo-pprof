package echopprof_test

import (
	echopprof "github.com/SkYNewZ/echo-pprof/v4"
	"github.com/labstack/echo/v4"
)

func ExampleWrap() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	echopprof.Wrap(e)

	e.Start(":8080")
}

func ExampleWrapGroup() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// echopprof also plays well with *echo.Group
	prefix := "/debug/pprof"
	group := e.Group(prefix)
	echopprof.WrapGroup(prefix, group)

	e.Start(":8080")
}
