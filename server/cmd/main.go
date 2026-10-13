package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Home")
	// })

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

	// Query Parameter
	mux.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintf(w, "Silahkan isi nama di query parameter")
		} else {
			fmt.Fprintf(w, "Hello %s", name)
		}
	})

	// Get Header
	mux.HandleFunc("/cek-header", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		contentType := r.Header.Get("Content-Type")

		authToken := r.Header.Get("X-Auth-Token")

		fmt.Fprintf(w, "User-Agent: %s\n", userAgent)
		fmt.Fprintf(w, "Content-Type: %s\n", contentType)
		fmt.Fprintf(w, "X-Auth-Token: %s\n\n", authToken)

		fmt.Fprintln(w, "--- SEMUA HEADER ---")
		for key, values := range r.Header {
			fmt.Fprintf(w, "%s: %v\n", key, values)
		}
	})

	// Form Post
	// Tampilan
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	// Endpoint untuk memproses data POST
	// mux.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodPost {
	// 		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	// 		return
	// 	}
	//
	// 	if err := r.ParseForm(); err != nil {
	// 		http.Error(w, "Gagal memparsing form", http.StatusBadRequest)
	// 		return
	// 	}
	//
	// 	nama := r.FormValue("nama")
	// 	email := r.FormValue("email")
	//
	// 	// Tampilan sederhana untuk hasil submit (bisa juga dibuatkan file HTML terpisah)
	// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 	fmt.Fprintf(w, `
	// 		<div style="font-family: sans-serif; text-align: center; margin-top: 50px;">
	// 			<h2 style="color: green;">Data Berhasil Diterima!</h2>
	// 			<p><b>Nama:</b> %s</p>
	// 			<p><b>Email:</b> %s</p>
	// 			<br><a href="/">Kembali ke Form</a>
	// 		</div>
	// 	`, nama, email)
	// })
// END: Form Post
	
// Response Code
	mux.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// 1. Cek Method (405 Method Not Allowed)
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Error: Gunakan method POST!")
			return
		}

		if err := r.ParseForm(); err != nil {
			// 2. Jika gagal parsing (400 Bad Request)
			http.Error(w, "Bad Request: Gagal memproses form", http.StatusBadRequest)
			return
		}

		nama := r.FormValue("nama")

		// 3. Validasi field kosong (422 Unprocessable Entity atau 400)
		if nama == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Validasi Gagal: Nama tidak boleh kosong!")
			return
		}

		// 4. Sukses menyimpan data (201 Created atau 200 OK)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Sukses! Data untuk %s berhasil dibuat.", nama)
	})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	fmt.Println("Server sedang berjalan di http://127.0.0.1:8080...")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
