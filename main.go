package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dbongo/app/db"
	"github.com/dbongo/app/middleware"
	"github.com/dbongo/app/router"
	"github.com/subosito/gotenv"
)

// load env vars
func init() {
	gotenv.Load(".env")
}

func main() {
	// setup database
	conn, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create the router and add middleware
	mux := router.New()
	mux.Use(middleware.Logger)
	http.Handle("/api/", mux)

	log.Printf("app started on port %s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
