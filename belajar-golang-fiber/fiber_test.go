package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/assert"
)

var engine = mustache.New("./template", ".mustache")

var app = fiber.New(fiber.Config{
	Views:        engine,
	IdleTimeout:  time.Second * 5,
	ReadTimeout:  time.Second * 5,
	WriteTimeout: time.Second * 5,
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString("Error " + err.Error())
	},
	Prefork: true,
})

func TestRouting(t *testing.T) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("TAI")
	})
	request := httptest.NewRequest("GET", "/", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "TAI", string(bytes))

}

func TestCtx(t *testing.T) {

	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "guests")
		return c.SendString("Hello " + name)
	})
	request := httptest.NewRequest("GET", "/hello?name=Nur", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Nur", string(bytes))

	request = httptest.NewRequest("GET", "/hello", nil)

	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err = io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello guests", string(bytes))

}

func TestHttpRequest(t *testing.T) {

	app.Get("/request", func(c *fiber.Ctx) error {
		first := c.Get("firstname")
		last := c.Cookies("lastname")
		return c.SendString("hello " + first + " " + last)
	})
	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "Nur")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Firdaus"})
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello Nur Firdaus", string(bytes))

}

func TestRouteParams(t *testing.T) {

	app.Get("/user/:userId/order/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Get Order " + orderId + " From User " + userId)
	})
	request := httptest.NewRequest("GET", "/user/1/order/2", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get Order 2 From User 1", string(bytes))

}

func TestFormRequest(t *testing.T) {

	app.Post("/createUser", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello " + name)
	})
	body := strings.NewReader("name=Nur")
	request := httptest.NewRequest("POST", "/createUser", body)
	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Nur", string(bytes))
}

//go:embed source/contoh.txt
var contohFile []byte

func TestFormUpload(t *testing.T) {

	app.Post("/uploadFile", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return c.SendString("upload success")
	})
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "contoh.txt")
	assert.Nil(t, err)

	file.Write(contohFile)
	writer.Close()
	request := httptest.NewRequest("POST", "/uploadFile", body)
	request.Header.Set("Content-type", writer.FormDataContentType())
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "upload success", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {

	app.Post("/Login", func(c *fiber.Ctx) error {
		body := c.Body()

		request := new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}

		return c.SendString("Hello " + request.Username)
	})
	body := strings.NewReader(`{"username":"Nur", "password":"rahasia"}`)
	request := httptest.NewRequest("POST", "/Login", body)
	request.Header.Set("Content-type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Nur", string(bytes))
}

func TestResponseJson(t *testing.T) {

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username": "Firdaus",
			"name":     "Nur Firdaus",
		})
	})
	request := httptest.NewRequest("GET", "/user", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"Nur Firdaus","username":"Firdaus"}`, string(bytes))
}

func TestDownloadFile(t *testing.T) {

	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./source/contoh.txt", "download.txt")
	})

	request := httptest.NewRequest("GET", "/download", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "attachment; filename=\"download.txt\"", response.Header.Get("Content-Disposition"))
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `this is sample file for uploads`, string(bytes))

}

func TestRoutingGroup(t *testing.T) {

	helloworld := func(c *fiber.Ctx) error {
		return c.SendString("Hello Nur")
	}

	api := app.Group("/api")
	api.Get("/product", helloworld)
	api.Get("/product/2", helloworld)

	web := app.Group("/web")
	web.Get("/user", helloworld)
	web.Get("/user/1", helloworld)

	request := httptest.NewRequest("GET", "/api/product/2", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `Hello Nur`, string(bytes))
}

func TestStatic(t *testing.T) {

	app.Static("/public", "./source")

	request := httptest.NewRequest("GET", "/public/contoh.txt", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `this is sample file for uploads`, string(bytes))
}

func TestError(t *testing.T) {

	app.Get("/error", func(c *fiber.Ctx) error {
		return errors.New("ups")
	})

	request := httptest.NewRequest("GET", "/error", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 500, response.StatusCode)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `Error ups`, string(bytes))
}

func TestView(t *testing.T) {
	app.Get("/view", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title":   "Hello Title",
			"header":  "Hello Header",
			"content": "Hello Content",
		})
	})

	request := httptest.NewRequest("GET", "/view", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), "Hello Title")
	assert.Contains(t, string(bytes), "Hello Header")
	assert.Contains(t, string(bytes), "Hello Content")
}

func TestHttpClient(t *testing.T) {
	client := fiber.AcquireClient()

	agent := client.Get("https://example.com/")
	status, response, errors := agent.String()

	assert.Nil(t, errors)
	assert.Equal(t, 200, status)
	assert.Contains(t, response, "Example Domain")
	defer fiber.ReleaseClient(client)
}
