package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "File is required")
		return
	}

	// preview file
	// http.ServeFile(w, r, "./resources/"+file)

	fileContent, err := os.ReadFile("./resources/" + file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
		return
	}
	w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Length", strconv.Itoa(len(fileContent)))
	w.Write(fileContent)
}

func TestDownloadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/download", DownloadFile)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
