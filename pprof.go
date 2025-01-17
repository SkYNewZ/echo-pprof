package echopprof

import (
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/labstack/echo/v5"
)

// Wrap adds several routes from package `net/http/pprof` to *echo.Echo object.
func Wrap(e *echo.Echo) {
	WrapGroup("", e.Group("/debug/pprof"))
}

// WrapGroup adds several routes from package `net/http/pprof` to *echo.Group object.
// Supported profiles from https://pkg.go.dev/runtime/pprof#Profile
func WrapGroup(prefix string, g *echo.Group) {
	routers := []struct {
		Method  string
		Path    string
		Handler echo.HandlerFunc
	}{
		{http.MethodGet, "", IndexHandler()},
		{http.MethodGet, "/", IndexHandler()},
		{http.MethodGet, "/goroutine", GoroutineHandler()},
		{http.MethodGet, "/heap", HeapHandler()},
		{http.MethodGet, "/allocs", AllocHandler()},
		{http.MethodGet, "/threadcreate", ThreadCreateHandler()},
		{http.MethodGet, "/block", BlockHandler()},
		{http.MethodGet, "/mutex", MutexHandler()},
		{http.MethodGet, "/cmdline", CmdlineHandler()},
		{http.MethodGet, "/profile", ProfileHandler()},
		{http.MethodGet, "/symbol", SymbolHandler()},
		{http.MethodPost, "/symbol", SymbolHandler()},
		{http.MethodGet, "/trace", TraceHandler()},
	}

	for _, r := range routers {
		g.Add(r.Method, strings.TrimPrefix(r.Path, prefix), r.Handler)
	}
}

// IndexHandler will pass the call from /debug/pprof to pprof.
func IndexHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Index(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// GoroutineHandler will pass the call from /debug/pprof/goroutine to pprof.
func GoroutineHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("goroutine").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// HeapHandler will pass the call from /debug/pprof/heap to pprof.
func HeapHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("heap").ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	}
}

// AllocHandler will pass the call from /debug/pprof/allocs to pprof.
func AllocHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("allocs").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// ThreadCreateHandler will pass the call from /debug/pprof/threadcreate to pprof.
func ThreadCreateHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// BlockHandler will pass the call from /debug/pprof/block to pprof.
func BlockHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("block").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// MutexHandler will pass the call from /debug/pprof/mutex to pprof.
func MutexHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("mutex").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// CmdlineHandler will pass the call from /debug/pprof/cmdline to pprof.
func CmdlineHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Cmdline(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// ProfileHandler will pass the call from /debug/pprof/profile to pprof.
func ProfileHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Profile(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// SymbolHandler will pass the call from /debug/pprof/symbol to pprof.
func SymbolHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Symbol(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// TraceHandler will pass the call from /debug/pprof/trace to pprof.
func TraceHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Trace(ctx.Response().Writer, ctx.Request())
		return nil
	}
}
