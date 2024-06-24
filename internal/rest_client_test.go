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
	for i := 0; i < 5; i++ {
		t.Log("test show request id=", i)
		err := run(testConfigFile, testHttpFile, showHandler(i))
		if err != nil {
			t.Error(err)
		}
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
