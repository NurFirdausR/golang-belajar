package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "Product" + id
		fmt.Fprintf(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1", string(body))
}

func TestNamedParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/item/:itemid", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemid := p.ByName("itemid")
		text := "Product " + id + " Item " + itemid
		fmt.Fprintf(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/item/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 2", string(body))
}

func TestCatchAllParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "Image " + image
		fmt.Fprintf(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/profile/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image /profile/profile.png", string(body))
}
