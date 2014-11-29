package middleware

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zenazn/goji/web"
)

const publicKey = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub

var verifyKey []byte

func init() {
	var err error
	verifyKey, err = ioutil.ReadFile(publicKey)
	if err != nil {
		log.Fatal("Error reading Private key")
		return
	}
}

// TokenAuth ...
func TokenAuth(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if err == nil && token.Valid {
			c.Env["token"] = token
			c.Env["jwt"] = token.Raw
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "missing access token\n")
		}
	}
	return http.HandlerFunc(fn)
}
