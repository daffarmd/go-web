package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HandlerParam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello %s", name)
}

func HandlerMultipleParam(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("firstname")
	lastName := r.URL.Query().Get("lastname")
	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func HandlerMultipleValueParam(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	nameMap := query["name"]
	fmt.Fprintf(w, "Hello %s", strings.Join(nameMap, " "))
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http://127.0.0.1:8080/hello/?name=Daffa", nil)
	recorder := httptest.NewRecorder()

	HandlerParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)

	fmt.Println(stringBody)

	assert.Equal(t, stringBody, "Hello Daffa")
}

func TestMultipleQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http://127.0.0.1:8080/hello/?firstname=Daffa&lastname=Ramadhan", nil)
	recorder := httptest.NewRecorder()

	HandlerMultipleParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)

	fmt.Println(stringBody)

	assert.Equal(t, stringBody, "Hello Daffa Ramadhan")
}

func TestMultipleValueParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http://127.0.0.1:8080/hello/?name=Daffa&name=Ramadhan", nil)
	recorder := httptest.NewRecorder()

	HandlerMultipleValueParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	stringBody := string(body)

	fmt.Println(stringBody)

	assert.Equal(t, stringBody, "Hello Daffa Ramadhan")
}
