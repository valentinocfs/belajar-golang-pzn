package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/master.gotmpl",
		"./templates/layout.gotmpl",
		"./templates/header.gotmpl",
		"./templates/footer.gotmpl",
	))
	t.ExecuteTemplate(w, "master.gotmpl", map[string]interface{}{
		"Title":   "Template Master",
		"Name":    "John Doe",
		"Content": "Hello, John Doe",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
