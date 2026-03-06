package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCodeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is required")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func TestResponseCode(t *testing.T) {
	requestName := "?name=John"
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello"+requestName, nil)
	recorder := httptest.NewRecorder()
	ResponseCodeHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", bodyString)
}
