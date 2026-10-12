package server

import (
	"fmt"
	"net/http"
	"testing"
)


func TestHandler(t *testing.T)  {	
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: handler,
	}
	
	fmt.Println("Server sedang berjalan di http://127.0.0.1:8080...")
	
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
