package main

import "github.com/subosito/gotenv"

// load env vars from .env so the app config can make use of them
func init() {
	gotenv.Load(".env")
}

// start app server
func main() {
	Serve()
}
