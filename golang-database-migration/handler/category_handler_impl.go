package handler

import (
	"golang-database-migration/helper"
	"golang-database-migration/model/web"
	"golang-database-migration/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryHandlerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &CategoryHandlerImpl{
		CategoryService: categoryService,
	}
}

func (handler *CategoryHandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var category web.CategoryCreateRequest
	helper.ReadFromRequestBody(request, &category)

	categoryResponse := handler.CategoryService.Save(request.Context(), category)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *CategoryHandlerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var category web.CategoryUpdateRequest
	helper.ReadFromRequestBody(request, &category)

	categoryId := params.ByName("categoryId")

	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	category.Id = categoryIdInt

	categoryResponse := handler.CategoryService.Update(request.Context(), category)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *CategoryHandlerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	handler.CategoryService.Delete(request.Context(), categoryIdInt)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *CategoryHandlerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	categoryIdInt, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := handler.CategoryService.FindById(request.Context(), categoryIdInt)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *CategoryHandlerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponse := handler.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
