package middleware

import (
	"log"
	"net/http"
	"time"
)

func MyLogger(request http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		request.ServeHTTP(w, r)
		log.Printf("Verb: %s | URL: %s | Time: %s | Bytes: %d", r.Method, r.URL.Path, startTime.Format(time.UnixDate), r.ContentLength)
	})
}