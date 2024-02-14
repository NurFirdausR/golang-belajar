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

func TestPanicHandler(t *testing.T) {

	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		fmt.Fprint(w, "Panic ", err)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// fmt.Fprintf(w, "Hello world")
		panic("ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic ups", string(body))
}
