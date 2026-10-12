package main

import (
	"fmt"
	"net/http"
)
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home")
	})

	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Product")
	})

	mux.HandleFunc("/resi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Resi")
	})

	mux.HandleFunc("/method-uri", func(w http.ResponseWriter, r *http.Request) {	
		methodURI := []any{r.Method, r.RequestURI}
		fmt.Fprint(w, methodURI)
	})

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	fmt.Println("Server sedang berjalan di http://127.0.0.1:8080...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
