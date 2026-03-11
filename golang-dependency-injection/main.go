package main

import (
	"fmt"
	"golang-dependency-injection/helper"
	"golang-dependency-injection/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(middleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    ":3000",
		Handler: middleware,
	}
}

func main() {
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Starting server at port 3000")
}
