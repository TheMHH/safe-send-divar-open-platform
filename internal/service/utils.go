package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	Token := os.Getenv("SECRET_KEY")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := ""
		if strings.HasPrefix(authHeader, "Bearer ") {
			bearerToken = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			log.Printf("No Bearer Token found")
			return
		}

		if bearerToken != Token {
			return
		}

		ctx := context.Background()
		r = r.WithContext(ctx)

		wrappedWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrappedWriter, r)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
