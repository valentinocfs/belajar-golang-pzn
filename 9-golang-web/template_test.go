package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := "<html><head><title>Template Golang</title></head><body>{{.}}</body></html>"
	// html, err = template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	html := template.Must(template.New("SIMPLE").Parse(templateText))

	html.ExecuteTemplate(w, "SIMPLE", "HTML TEMPLATE")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	html.ExecuteTemplate(w, "simple.gohtml", "HTML TEMPLATE GOLANG")
}

func TestTemplateFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseGlob("./templates/*.gohtml"))
	html.ExecuteTemplate(w, "simple.gohtml", "HTML TEMPLATE GLOB GOLANG ")
}

func TestTemplateFileGlob(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	html.ExecuteTemplate(w, "simple.gohtml", "HTML TEMPLATE EMBED GOLANG ")
}

func TestTemplateFileEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
