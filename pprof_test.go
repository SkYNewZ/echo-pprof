package echopprof

import (
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v5"
)

func newServer() *echo.Echo {
	e := echo.New()
	return e
}

func checkRouters(routes []echo.RouteInfo, t *testing.T, expectedRouters map[string][]string) {
	for _, router := range routes {
		if (router.Method() != http.MethodGet && router.Method() != http.MethodPost) || strings.HasSuffix(router.Path(), "/*") {
			continue
		}

		names, ok := expectedRouters[router.Path()]
		if !ok {
			t.Errorf("missing router %s", router.Path())
		}

		for _, name := range names {
			if !strings.Contains(router.Name(), name) {
				t.Errorf("handler for %s should contains %s, got %s", router.Path(), name, router.Name())
			}
		}
	}
}

// go test github.com/SkYNewZ/echo-pprof/v5 -v -run=TestWrap\$
func TestWrap(t *testing.T) {
	e := newServer()
	Wrap(e)

	// this is now the default route name
	// See https://github.com/labstack/echo/discussions/2000
	expectedRouters := map[string][]string{
		"/debug/pprof":              {"GET:/debug/pprof"},
		"/debug/pprof/":             {"GET:/debug/pprof/"},
		"/debug/pprof/goroutine":    {"GET:/debug/pprof/goroutine"},
		"/debug/pprof/heap":         {"GET:/debug/pprof/heap"},
		"/debug/pprof/allocs":       {"GET:/debug/pprof/allocs"},
		"/debug/pprof/threadcreate": {"GET:/debug/pprof/threadcreate"},
		"/debug/pprof/block":        {"GET:/debug/pprof/block"},
		"/debug/pprof/mutex":        {"GET:/debug/pprof/mutex"},
		"/debug/pprof/cmdline":      {"GET:/debug/pprof/cmdline"},
		"/debug/pprof/profile":      {"GET:/debug/pprof/profile"},
		"/debug/pprof/symbol":       {"/debug/pprof/symbol"},
		"/debug/pprof/trace":        {"GET:/debug/pprof/trace"},
	}

	checkRouters(e.Router().Routes(), t, expectedRouters)
}

// go test github.com/SkYNewZ/echo-pprof/v5 -v -run=TestWrapGroup\$
func TestWrapGroup(t *testing.T) {
	for _, prefix := range []string{"/debug"} {
		e := newServer()
		g := e.Group(prefix)
		WrapGroup(prefix, g)
		baseRouters := map[string][]string{
			"":              {"GET:/debug"},
			"/":             {"GET:/debug/"},
			"/goroutine":    {"GET:/debug/goroutine"},
			"/heap":         {"GET:/debug/heap"},
			"/allocs":       {"GET:/debug/allocs"},
			"/threadcreate": {"GET:/debug/threadcreate"},
			"/block":        {"GET:/debug/block"},
			"/mutex":        {"GET:/debug/mutex"},
			"/cmdline":      {"GET:/debug/cmdline"},
			"/profile":      {"GET:/debug/profile"},
			"/symbol":       {"/debug/symbol"},
			"/trace":        {"GET:/debug/trace"},
		}

		expectedRouters := make(map[string][]string, len(baseRouters))
		for r, h := range baseRouters {
			expectedRouters[prefix+r] = h
		}

		checkRouters(e.Router().Routes(), t, expectedRouters)
	}
}
