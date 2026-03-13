package main

import (
	"fmt"
	"golang-database-migration/app"
	"golang-database-migration/handler"
	"golang-database-migration/helper"
	"golang-database-migration/middleware"
	"golang-database-migration/repository"
	"golang-database-migration/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router := app.NewRouter(categoryHandler)
	middleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    ":3000",
		Handler: middleware,
	}

	fmt.Println("Starting server at port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
