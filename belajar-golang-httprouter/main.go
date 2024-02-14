package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Helo cuy")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
