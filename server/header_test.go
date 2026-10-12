package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprint(w, contentType)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://127.0.0.1:8080/header", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	GetHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)

	fmt.Println(stringBody)

	assert.Equal(t, stringBody, "application/json")
}


func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "Daffa")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://127.0.0.1:8080/header", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	powerBy := recorder.Header().Get("x-powered-by")
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)
	fmt.Println(powerBy)
	fmt.Println(stringBody)

	assert.Equal(t, powerBy, "Daffa")
	assert.Equal(t, stringBody, "OK")

}
