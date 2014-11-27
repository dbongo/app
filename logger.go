package main

import (
	"log"
	"net/http"
	"time"

	"github.com/zenazn/goji/web"
)

// StatusResponseWriter ...
type StatusResponseWriter struct {
	http.ResponseWriter
	Status int
}

// WriteHeader ...
func (w *StatusResponseWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

// Logger ...
func Logger(c *web.C, h http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &StatusResponseWriter{w, http.StatusOK}
		h.ServeHTTP(sw, r)
		var remoteAddr string
		fwd := r.Header.Get("X-Forwarded-For")
		if fwd == "" {
			remoteAddr = r.RemoteAddr
		} else {
			remoteAddr = fwd + ":" + r.Header.Get("X-Forwarded-Port")
		}
		log.Printf("%s %6.4fms %s %d %s\n", remoteAddr, float64(time.Since(start))/float64(time.Millisecond), r.Method, sw.Status, r.RequestURI)
	}
	return http.HandlerFunc(handler)
}
