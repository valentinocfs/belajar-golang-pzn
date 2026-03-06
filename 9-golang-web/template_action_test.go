package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gotmpl"))
	t.ExecuteTemplate(w, "if.gotmpl", map[string]interface{}{
		"Title": "Template Action If",
		// "Name":  "John Doe",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionComparator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gotmpl"))
	t.ExecuteTemplate(w, "comparator.gotmpl", map[string]interface{}{
		"Title":      "Template Action Comparator",
		"FinalValue": 80,
	})
}

func TestTemplateActionComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gotmpl"))
	t.ExecuteTemplate(w, "range.gotmpl", map[string]interface{}{
		"Title": "Template Action Range",
		"Value": []string{"John", "Jane", "Ilham"},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gotmpl"))
	t.ExecuteTemplate(w, "with.gotmpl", map[string]interface{}{
		"Title": "Template Action With",
		"Name":  "John Doe",
		"Address": map[string]interface{}{
			"Street": "123 Main St",
			"City":   "Anytown",
			"State":  "CA",
			"Zip":    "12345",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
