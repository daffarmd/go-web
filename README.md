# Go Web

Menggunakan standard library golang yaitu `net/http` tanpa bantuan framework pihak ke-3.
Rata-Rata User Golang jarang menggunakan framework karena alasan flexibelitas yang tinggi.

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

## 5. HTTP Test

Berfungsi untuk menjalankan HTTP secara unit test tanpa kita repot untuk menyalakan server :
`httptest.NewRequest()` ini untuk request

`httptest.NewRecorder()` ini untuk response

```go
func HelloHandler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprint(w, "Hello Handler")
}

func TestHelloHandler(t *testing.T) {
 request := httptest.NewRequest("GET", "http://127.0.0.1:8080/hello", nil)
 recorder := httptest.NewRecorder()

 HelloHandler(recorder, request)

 response := recorder.Result()
 body, _ := io.ReadAll(response.Body)
 stringBody := string(body)

 assert.Equal(t, stringBody, "Hello Handler")
}
```

## 6. QUERY Parameter

Untuk mengirimkan sebuah data pada parameter kita bisa menggunakan

Set Parameter nya :

```go
 r.URL.Query().Get("nama_parameter")
```

Cara menggunakan nya :

```go
request := httptest.NewRequest("GET", "http://127.0.0.1:8080/hello/?nama_parameter=Daffa", nil)
```

Contoh != Unit Test :

`@server/cmd/main.go` = Query Parameter

## 7. Header

Kita bisa mengambil header yang tersedia pada service http dengan cara :

Set di sisi server :

```go
w.Header().Add("x-powered-by", "Daffa")
```

Client Mengambil Header Spesifik :

```go
powerBy := recorder.Header().Get("x-powered-by")
```

Contoh != Unit Test :

`@server/cmd/main.go` = Get Header

## 8. Form Post
Selain kita bisa mengirimkan data melalui Query Parameter kita juga bisa mengirimkan sebuah data menggunakan Form Post dengan cara :

Pada Handler :
```go
firstName := r.PostForm.Get("first_name")
lastName := r.PostForm.Get("last_name")
```

Pada Unit Test :

```go
// 1. Menentukan Body atau request yang akan dikirimkan 
requestBody := strings.NewReader("first_name=Daffa&last_name=Ramadhan")
request := httptest.NewRequest("POST", "http://127.0.0.1:8088", requestBody)
// 2. Jika Form Post maka Type nya adalah application/x-www-form-urlencoded
request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
```

Contoh != Unit Test :

`@server/cmd/main.go` = Form Post

## 9. Response Code
Setiap kita melakukan request dan response maka server akan mengembalikan status berupa code yang di sebut Response Code status ini bertujuan memberli tahu client apakah 
Proses yang di lakukan berhasil atau gagal :
`200` = Success
`404` = Not Found

Untuk lengkapnya bisa di cek di `https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status`

Contoh Kode :

```go
func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", "http://127.0.0.1:8080/?name=Daffa", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	responseCode := response.StatusCode
	status := response.Status
	stringBody := string(body)
	fmt.Println(responseCode) // 200
	fmt.Println(status) // 200 OK
	fmt.Println(stringBody) // Hello Daffa
	assert.Equal(t, stringBody, "Hello Daffa")
}
```

Contoh Impelementasi (Bukan Unit Test) :
`@server/cmd/main.go` = Response Code

