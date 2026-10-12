package server

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest (t *testing.T) {
	var handler http.HandlerFunc = func (w http.ResponseWriter, r *http.Request)  {
		method := r.Method
		fmt.Fprintf(w, "Method : %s", method)
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
