package main

import (
	"log"
	"net/http"

	"github.com/dbongo/app/conf"
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
	// create app config
	conf, err := conf.New()
	if err != nil {
		log.Fatal(err)
	}

	// setup database
	db, err := db.Open(conf.DBHost, conf.DBName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("app connected to %s at %s", conf.DBName, conf.DBHost)
	defer db.Close()

	// create the router and add middleware
	mux := router.New()
	mux.Use(middleware.Logger)
	http.Handle("/api/", mux)

	log.Printf("app started on port %s", conf.Port)
	log.Fatal(http.ListenAndServe(":"+conf.Port, nil))
}
