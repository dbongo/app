package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

// createRouter returns app router
func createRouter() *web.Mux {
	mux := web.New()
	mux.Get("/api/hello", helloWorld)
	mux.Get("/api/hello/:name", helloName)
	return mux
}

func helloWorld(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func helloName(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}
