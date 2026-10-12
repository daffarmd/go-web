# Go Web

## 1. Cara Jalankan Server

```go
server := http.Server{
 Addr: "127.0.0.1:8080",
}
```

## 2. Cara Menambahkan Handler

Handler?

Handler di Golang adalah sebuah fungsi atau objek yang bertugas menerima dan merespons request HTTP yang masuk dari klien (seperti browser atau aplikasi mobile).

```go
var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
 fmt.Fprint(w, "Hello World")
}
```

```go
server := http.Server{
 Addr:    "127.0.0.1:8080",
 Handler: handler,
}
```

## 3. Cara Routing

Untuk Cara Routing bisa menggunakan `http.NewServeMux()`

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 fmt.Fprint(w, "Home")
})
```

```go
mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
 fmt.Fprint(w, "Product")
})
```

```go
server := http.Server{
 Addr:    "127.0.0.1:8080",
 Handler: mux,
}
```

## 4. Request

struct yang merepresentasikan HTTP Request yang di kirim oleh Web Browser
semua informasi yang di kirim bisa kita dapatkan di Request
Seperti, URL, http Method, http Header, http body, dll
Dengan Menggunakan `r.Method`

```go
var handler http.HandlerFunc = func (w http.ResponseWriter, r *http.Request){
method := r.Method
fmt.Fprintf(w, "Method : %s", method)
}
```
