package server

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "127.0.0.1:8088",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
