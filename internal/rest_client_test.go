package internal

import (
	"context"
	"net/http"
	"testing"
	"time"
)

var (
	testConfigFile = "./testdata/conf.yaml"
	testHttpFile   = "./testdata/request.http"
)

func TestShowRequest(t *testing.T) {
	err := run(testConfigFile, testHttpFile, showHandler(0))
	if err != nil {
		t.Error(err)
	}
}

func TestSendRequest(t *testing.T) {
	shutdown := startTestServer()
	defer shutdown()
	err := run(testConfigFile, testHttpFile, sendHandler(1))
	if err != nil {
		t.Error(err)
	}
}

func startTestServer() func() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /info", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is a test server"))
	})

	mux.HandleFunc("POST /say", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	server := &http.Server{Addr: ":8080", Handler: mux}
	go server.ListenAndServe()
	time.Sleep(time.Second)
	return func() {
		server.Shutdown(context.Background())
	}
}
