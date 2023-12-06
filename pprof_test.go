package echopprof

import (
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func newServer() *echo.Echo {
	e := echo.New()
	return e
}

func checkRouters(routers []*echo.Route, t *testing.T, expectedRouters map[string]string) {
	for _, router := range routers {
		if (router.Method != "GET" && router.Method != "POST") || strings.HasSuffix(router.Path, "/*") {
			continue
		}
		name, ok := expectedRouters[router.Path]
		if !ok {
			t.Errorf("missing router %s", router.Path)
		}
		if !strings.Contains(router.Name, name) {
			t.Errorf("handler for %s should contain %s, got %s", router.Path, name, router.Name)
		}
	}
}

// go test github.com/SkYNewZ/echo-pprof/v4 -v -run=TestWrap\$
func TestWrap(t *testing.T) {
	e := newServer()
	Wrap(e)

	expectedRouters := map[string]string{
		"/debug/pprof":              "IndexHandler",
		"/debug/pprof/":             "IndexHandler",
		"/debug/pprof/goroutine":    "GoroutineHandler",
		"/debug/pprof/heap":         "HeapHandler",
		"/debug/pprof/allocs":       "AllocHandler",
		"/debug/pprof/threadcreate": "ThreadCreateHandler",
		"/debug/pprof/block":        "BlockHandler",
		"/debug/pprof/mutex":        "MutexHandler",
		"/debug/pprof/cmdline":      "CmdlineHandler",
		"/debug/pprof/profile":      "ProfileHandler",
		"/debug/pprof/symbol":       "SymbolHandler",
		"/debug/pprof/trace":        "TraceHandler",
	}

	checkRouters(e.Routes(), t, expectedRouters)
}

// go test github.com/SkYNewZ/echo-pprof/v4 -v -run=TestWrapGroup\$
func TestWrapGroup(t *testing.T) {
	for _, prefix := range []string{"/debug"} {
		e := newServer()
		g := e.Group(prefix)
		WrapGroup(prefix, g)
		baseRouters := map[string]string{
			"":              "IndexHandler",
			"/":             "IndexHandler",
			"/goroutine":    "GoroutineHandler",
			"/heap":         "HeapHandler",
			"/allocs":       "AllocHandler",
			"/threadcreate": "ThreadCreateHandler",
			"/block":        "BlockHandler",
			"/mutex":        "MutexHandler",
			"/cmdline":      "CmdlineHandler",
			"/profile":      "ProfileHandler",
			"/symbol":       "SymbolHandler",
			"/trace":        "TraceHandler",
		}
		expectedRouters := make(map[string]string, len(baseRouters))
		for r, h := range baseRouters {
			expectedRouters[prefix+r] = h
		}
		checkRouters(e.Routes(), t, expectedRouters)
	}
}
