package hashivault

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var testServer *httptest.Server
var mux = &sync.Mutex{}
var handlers = map[string]http.HandlerFunc{}
var loginCount = 0

type closerFunc func()

func startTestServer(t *testing.T) (string, *http.Client, closerFunc) {
	t.Helper()

	mux.Lock()
	defer mux.Unlock()

	closer := func() {
		mux.Lock()
		defer mux.Unlock()
		testServer.Close()
	}

	if testServer != nil {
		return testServer.URL, testServer.Client(), closer
	}

	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.Lock()
		defer mux.Unlock()

		p := r.URL.String()
		if h, ok := handlers[p]; ok {
			h(w, r)
			return
		}
		t.Errorf("no handler for %s", p)
		w.WriteHeader(http.StatusNotFound)
	}))

	loginHandler := func(w http.ResponseWriter, r *http.Request) {
		loginCount++
		w.WriteHeader(http.StatusOK)
		tokenResponse := fmt.Sprintf(ghVaultResponseTemplate, fmt.Sprintf("token-%d", loginCount))
		fmt.Fprintln(w, tokenResponse)
	}
	handlers["/v1/auth/github/login"] = loginHandler

	return testServer.URL, testServer.Client(), closer
}

func addHandler(path string, handler http.HandlerFunc) {
	mux.Lock()
	defer mux.Unlock()
	handlers[path] = handler
}
