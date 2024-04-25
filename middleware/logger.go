package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging wraps the passed through HTTP handler and logs on completion of that
// handler the following information:
//
//	{StatusCode} {RequestMethod} {RequestPath} {TimeTaken}
//	        200             GET        /gophi    18.091µs
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lw := &logWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(lw, r)

		log.Println(lw.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

type logWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader overrides the embedded http.ResponseWriter calls to
// (Responsewriter).WriteHeader() and captures the returned status
// code to the caller. This is then set on the wrapped writer to
// be logged out by the logger middleware after the request completes.
func (w *logWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}
