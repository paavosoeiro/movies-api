package middleware

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func newResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           &bytes.Buffer{},
	}
}

func (rw *customResponseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *customResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)

		log.Printf("Starting %s %s", r.Method, r.RequestURI)

		next.ServeHTTP(rw, r)

		log.Printf("Finishing %s %s with status %d in %v", r.Method, r.RequestURI, rw.statusCode, time.Since(start))
		log.Printf("Response Body %s", rw.body.String())
	})
}
