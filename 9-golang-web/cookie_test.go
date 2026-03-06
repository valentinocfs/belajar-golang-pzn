package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = "1234567890"
	cookie.MaxAge = 3600
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Cookie has been set")
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")

	if err != nil {
		fmt.Fprint(w, "Cookie not found")
		return
	} else {
		fmt.Fprintf(w, "%s : %s", cookie.Name, cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookieHandler)
	mux.HandleFunc("/get-cookie", GetCookieHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get-cookie", nil)
	request.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: "1234567890",
	})
	response := httptest.NewRecorder()
	GetCookieHandler(response, request)

	responseResult := response.Result()
	body, _ := io.ReadAll(responseResult.Body)
	bodyString := string(body)

	fmt.Println("Response Status:", responseResult.Status)
	fmt.Println("Response Body:", bodyString)
}
