package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

var engine = mustache.New("./template", ".mustache")

var app = fiber.New(fiber.Config{
	Views: engine,
	ErrorHandler: func(ctx fiber.Ctx, err error) error {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString("Error: " + err.Error())
	},
})

//go:embed source/sample.txt
var sampleFile []byte

func TestRouting(t *testing.T) {
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, World!", string(bytes))
}

func TestContext(t *testing.T) {
	app.Get("/hello", func(ctx fiber.Ctx) error {
		name := ctx.Query("name", "Guest")
		return ctx.SendString("Hello, " + name)
	})

	request := httptest.NewRequest(http.MethodGet, "/hello?name=John", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, John", string(bytes))
	fmt.Println(string(bytes))

	request = httptest.NewRequest(http.MethodGet, "/hello", nil)
	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, Guest", string(bytes))
	fmt.Println(string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(ctx fiber.Ctx) error {
		firstname := ctx.Get("firstname")
		lastname := ctx.Cookies("lastname")
		return ctx.SendString("Hello " + firstname + " " + lastname)
	})

	request := httptest.NewRequest(http.MethodGet, "/request", nil)
	request.Header.Set("firstname", "John")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Doe"})

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello John Doe", string(bytes))
	fmt.Println(string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(ctx fiber.Ctx) error {
		userId := ctx.Params("userId")
		orderId := ctx.Params("orderId")
		return ctx.SendString("User: " + userId + ", Order: " + orderId)
	})

	request := httptest.NewRequest(http.MethodGet, "/users/1/orders/456", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "User: 1, Order: 456", string(bytes))
	fmt.Println(string(bytes))
}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(ctx fiber.Ctx) error {
		name := ctx.FormValue("name")
		return ctx.SendString("Hello " + name)
	})

	body := strings.NewReader("name=John")
	request := httptest.NewRequest(http.MethodPost, "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello John", string(bytes))
	fmt.Println(string(bytes))
}

func TestMultipartFormRequest(t *testing.T) {
	app.Post("/upload", func(ctx fiber.Ctx) error {
		file, err := ctx.FormFile("file")
		if err != nil {
			return err
		}

		err = ctx.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return ctx.SendString(file.Filename + " uploaded successfully")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "sample.txt")
	assert.Nil(t, err)
	file.Write(sampleFile)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "sample.txt uploaded successfully", string(bytes))
	fmt.Println(string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(ctx fiber.Ctx) error {
		body := ctx.Body()

		request := new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}

		return ctx.SendString("Hello " + request.Username)
	})

	body := strings.NewReader(`{"username": "johncena", "password": "secret"}`)
	request := httptest.NewRequest(http.MethodPost, "/login", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello johncena", string(bytes))
	fmt.Println(string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(ctx fiber.Ctx) error {
		request := new(RegisterRequest)
		err := ctx.Bind().Body(request)
		if err != nil {
			return err
		}

		return ctx.SendString("Register Success with username " + request.Username)
	})
}

func TestBodyParserJSON(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`{"username": "johncena", "password": "secret", "name": "John Cena"}`)
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success with username johncena", string(bytes))
	fmt.Println(string(bytes))
}

func TestBodyParserForm(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`username=johncena&password=secret&name=John Cena`)
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success with username johncena", string(bytes))
	fmt.Println(string(bytes))
}

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`
		<RegisterRequest>
			<username>johncena</username>
			<password>secret</password>
			<name>John Cena</name>
		</RegisterRequest>`)

	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/xml")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success with username johncena", string(bytes))
	fmt.Println(string(bytes))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"username": "johncena",
			"name":     "John Cena",
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/user", nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"John Cena","username":"johncena"}`, string(bytes))
	fmt.Println(string(bytes))
}

func TestResponseXML(t *testing.T) {
	type SomeStruct struct {
		XMLName xml.Name `xml:"User"`
		Name    string   `xml:"Name"`
		Age     uint8    `xml:"Age"`
	}

	app.Get("/user", func(ctx fiber.Ctx) error {
		return ctx.XML(SomeStruct{
			XMLName: xml.Name{Local: "User"},
			Name:    "John Cena",
			Age:     30,
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/user", nil)
	request.Header.Set("Content-Type", "application/xml")

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `<User><Name>John Cena</Name><Age>30</Age></User>`, string(bytes))
	fmt.Println(string(bytes))
}

func TestDownloadFile(t *testing.T) {
	app.Get("/download", func(ctx fiber.Ctx) error {
		return ctx.Download("./source/sample.txt", "sample.txt")
	})

	request := httptest.NewRequest(http.MethodGet, "/download", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "attachment; filename=\"sample.txt\"", response.Header.Get("Content-Disposition"))

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "This is a sample text file.\r\n", string(bytes))
	fmt.Println(string(bytes))
}

func TestRoutingGroup(t *testing.T) {
	helloWorld := func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello World")
	}

	api := app.Group("/api")
	api.Get("/hello", helloWorld)
	api.Get("/world", helloWorld)

	web := app.Group("/web")
	web.Get("/hello", helloWorld)
	web.Get("/world", helloWorld)

	request := httptest.NewRequest(http.MethodGet, "/api/hello", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
	fmt.Println(string(bytes))
}

func TestStatic(t *testing.T) {
	app.Use("/public", static.New("", static.Config{
		FS:     os.DirFS("source"),
		Browse: true,
	}))

	request := httptest.NewRequest(http.MethodGet, "/public/sample.txt", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "This is a sample text file.\r\n", string(bytes))
	fmt.Println(string(bytes))
}

func TestErrorHandling(t *testing.T) {
	app.Get("/error", func(ctx fiber.Ctx) error {
		return errors.New("something went wrong")
	})

	request := httptest.NewRequest(http.MethodGet, "/error", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Error: something went wrong", string(bytes))
	fmt.Println(string(bytes))
}

func TestTemplate(t *testing.T) {
	app.Get("/view", func(ctx fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"title":   "This is Title",
			"header":  "This is Header",
			"content": "This is Content",
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/view", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "This is Title")
	assert.Contains(t, string(bytes), "This is Header")
	assert.Contains(t, string(bytes), "This is Content")
	fmt.Println(string(bytes))
}

func TestClient(t *testing.T) {
	client := client.New().SetBaseURL("https://example.com/")

	resp, err := client.Get("/")
	if err != nil {
		panic(err)
	}
	defer resp.Close()

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, resp.String(), "Example Domain")
}
