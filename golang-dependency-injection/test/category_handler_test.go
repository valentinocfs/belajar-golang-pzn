package test

import (
	"database/sql"
	"encoding/json"
	"golang-dependency-injection/app"
	"golang-dependency-injection/handler"
	"golang-dependency-injection/helper"
	"golang-dependency-injection/middleware"
	"golang-dependency-injection/repository"
	"golang-dependency-injection/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func setupDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang_restful_api_test?parseTime=True")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router := app.NewRouter(categoryHandler)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	_, err := db.Exec("TRUNCATE category")
	helper.PanicIfError(err)
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Fashion"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 200, int(responseMap["code"].(float64)))
	assert.Equal(t, "OK", responseMap["status"])
	assert.Equal(t, "Fashion", responseMap["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, 400, int(responseMap["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseMap["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Electronics"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/1", requestBody)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 200, int(responseMap["code"].(float64)))
	assert.Equal(t, "OK", responseMap["status"])
	assert.Equal(t, "Electronics", responseMap["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Electronics"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/1", requestBody)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, 404, int(responseMap["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseMap["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 200, int(responseMap["code"].(float64)))
	assert.Equal(t, "OK", responseMap["status"])
	assert.Equal(t, "Electronics", responseMap["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupDB()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/100", nil)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")

	recorder := httptest.NewRecorder()

	setupRouter(db).ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, 404, int(responseMap["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseMap["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupDB()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/1", nil)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")

	recorder := httptest.NewRecorder()

	setupRouter(db).ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 200, int(responseMap["code"].(float64)))
	assert.Equal(t, "OK", responseMap["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupDB()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/100", nil)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")

	recorder := httptest.NewRecorder()

	setupRouter(db).ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, 404, int(responseMap["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseMap["status"])
}

func TestListCategorySuccess(t *testing.T) {
	db := setupDB()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "SECRET_API_KEY")

	recorder := httptest.NewRecorder()

	setupRouter(db).ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 200, int(responseMap["code"].(float64)))
	assert.Equal(t, "OK", responseMap["status"])
	assert.Equal(t, 0, len(responseMap["data"].([]interface{})))
}

func TestUnauthorized(t *testing.T) {
	db := setupDB()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "WRONG_SECRET_API_KEY")

	recorder := httptest.NewRecorder()

	setupRouter(db).ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseMap map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseMap)

	assert.Equal(t, 401, int(responseMap["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseMap["status"])
}
