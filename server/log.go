package server

import (
	"log"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Now().Sub(start)
		log.Printf("Request %s %s %s took %v", r.Method, r.URL.Path, r.Proto, duration)
	}
}
