package app

import (
	"golang-dependency-injection/exception"
	"golang-dependency-injection/handler"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryHandler handler.CategoryHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryHandler.FindAll)
	router.GET("/api/categories/:categoryId", categoryHandler.FindById)
	router.POST("/api/categories", categoryHandler.Create)
	router.PUT("/api/categories/:categoryId", categoryHandler.Update)
	router.DELETE("/api/categories/:categoryId", categoryHandler.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
