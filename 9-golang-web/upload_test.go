package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var myNewTemplates = template.Must(template.ParseGlob("templates/*.gotmpl"))

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myNewTemplates.ExecuteTemplate(w, "upload.form.gotmpl", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(32 << 20) // 32MB default size
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")

	myNewTemplates.ExecuteTemplate(w, "upload.success.gotmpl", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/icon.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "John Doe")
	file, _ := writer.CreateFormFile("file", "icon.png")
	file.Write(uploadFileTest)

	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
	bodyResult, _ := io.ReadAll(response.Body)
	bodyString := string(bodyResult)

	fmt.Println(bodyString)
}
