package main

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2"
)

type app struct {
	conf   *config
	db     *mongodb
	router *web.Mux
}

// createApp returns a confirgured app
func createApp() (*app, error) {
	var err error
	app := &app{}

	// set the app config
	app.conf, err = createConfig()
	if err != nil {
		return app, err
	}

	// set and initialize the app db
	app.db = createDB()
	if err := app.initDB(); err != nil {
		return app, err
	}

	// setup app router and add middleware
	app.router = createRouter()
	app.router.Use(Logger)

	http.Handle("/api/", app.router)
	return app, nil
}

// close removes the mongo session for app
func (app *app) close() {
	app.db.session.Close()
}

// initDB creates mongo session for app
func (app *app) initDB() error {
	session, err := mgo.Dial(app.conf.DBHost)
	if err != nil {
		return err
	}
	log.Printf("app connected to mongodb at %s", app.conf.DBHost)
	session.SetMode(mgo.Monotonic, true)
	app.db.session = session
	return nil
}
