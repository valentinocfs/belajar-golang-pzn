package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeaderHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprintf(w, "Content Type: %s", contentType)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	RequestHeaderHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func ResponseHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Golang")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()
	ResponseHeaderHandler(recorder, request)

	response := recorder.Result()
	xPoweredBy := response.Header.Get("X-Powered-By")
	fmt.Println("X-Powered-By:", xPoweredBy)
}
