package service

import (
	"net/http"
	"os"
)

type Server struct {
	hmacSecret string
}

func NewServer() *Server {
	return &Server{
		hmacSecret: os.Getenv("SECRET_KEY"),
	}
}

func Init(s *Server) {
	http.Handle("/url1", Middleware(http.HandlerFunc(s.func1)))
	http.Handle("/url2", Middleware(http.HandlerFunc(s.func2)))
	http.Handle("/url3", Middleware(http.HandlerFunc(s.func3)))
}
