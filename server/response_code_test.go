package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	fmt.Println(responseCode)
	fmt.Println(status)
	fmt.Println(stringBody)
	assert.Equal(t, stringBody, "Hello Daffa")
}
