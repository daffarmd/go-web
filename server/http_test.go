package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
