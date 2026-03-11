//go:build wireinject
// +build wireinject

package main

import (
	"golang-dependency-injection/app"
	"golang-dependency-injection/handler"
	"golang-dependency-injection/middleware"
	"golang-dependency-injection/repository"
	"golang-dependency-injection/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	handler.NewCategoryHandler,
	wire.Bind(new(handler.CategoryHandler), new(*handler.CategoryHandlerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		NewValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

func NewValidator() *validator.Validate {
	return validator.New()
}
