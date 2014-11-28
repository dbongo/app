package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"code.google.com/p/go.net/context"

	webcontext "github.com/goji/context"
	"github.com/zenazn/goji/web"

	"github.com/dbongo/app/middleware"
	"github.com/dbongo/app/model"
	"github.com/dbongo/app/router"
	"github.com/subosito/gotenv"
)

// location of the files used for signing and verification of jwt token
const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa 1024
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey, signKey []byte
)

func init() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
	}

	gotenv.Load(".env")
}

func main() {
	// create the router and add middleware
	mux := router.New()
	mux.Use(middleware.Options)
	mux.Use(ContextMiddleware)
	mux.Use(middleware.Logger)
	http.Handle("/api/", mux)

	user := model.User{Username: "bob", Password: "bob123"}
	err := user.Create()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("app listenining on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// ContextMiddleware creates a new go.net/context and injects into the current goji context.
func ContextMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		webcontext.Set(c, ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
