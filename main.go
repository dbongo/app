package main

import (
	"io/ioutil"
	"log"
	"net/http"

	webcontext "github.com/goji/context"

	"github.com/dbongo/app/db"
	"github.com/dbongo/app/middleware"
	"github.com/dbongo/app/model"
	"github.com/dbongo/app/router"
	"github.com/subosito/gotenv"
)

const privateKey = "keys/app.rsa" // openssl genrsa -out app.rsa 1024

var signKey []byte

func init() {
	var err error
	signKey, err = ioutil.ReadFile(privateKey)
	if err != nil {
		log.Fatal("Error reading private key")
	}
	gotenv.Load(".env")
}

func main() {
	// create the router and add middleware
	mux := router.New()
	mux.Use(webcontext.Middleware)
	mux.Use(middleware.Options)
	mux.Use(middleware.Logger)
	mux.Use(middleware.TokenAuth)
	http.Handle("/api/", mux)

	sess, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()
	err = sess.DB().DropDatabase()
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{Email: "bob@email.com", Username: "bob", Password: "bob123"}
	err = user.Create()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("app listenining on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
