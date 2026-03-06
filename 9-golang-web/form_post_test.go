package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPostHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Fprint(writer, "Error:", err.Error())
		return
	}
	firstname := request.PostForm.Get("firstname")
	lastname := request.PostForm.Get("lastname")

	// Atau bisa dgn seperti ini tanpa parsing
	// firstname := request.PostFormValue("firstname")
	// lastname := request.PostFormValue("lastname")

	fmt.Fprintf(writer, "Hello %s %s!", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=John&lastname=Doe")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/hello", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPostHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
