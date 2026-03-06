package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", bodyString)
}

func SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello World")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func TestQueryParameter(t *testing.T) {
	// request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=John", nil)
	recorder := httptest.NewRecorder()

	SayHelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", bodyString)
}

func MultipleQueryParameterHandler(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")
	fmt.Fprintf(w, "\nFirst Name: %s, Last Name: %s", firstname, lastname)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?firstname=John&lastname=Doe", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameterHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", bodyString)
}

func MultipleValueQueryParameterHandler(w http.ResponseWriter, r *http.Request) {
	// ?name=John&name=Jane&name=Bob
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprintf(w, "Names: %v", strings.Join(names, ", "))
}

func TestMultipleValueQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=John&name=Jane&name=Bob", nil)
	recorder := httptest.NewRecorder()

	MultipleValueQueryParameterHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", bodyString)
}
