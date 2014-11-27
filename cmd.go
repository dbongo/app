package main

import (
	"log"
	"net/http"
)

// Serve runs the HTTP server
func Serve() {
	app, err := createApp()
	if err != nil {
		log.Fatal(err)
	}
	defer app.close()
	log.Printf("app started on port %s", app.conf.ServerPort)
	log.Fatal(http.ListenAndServe(app.conf.ServerPort, nil))
}
