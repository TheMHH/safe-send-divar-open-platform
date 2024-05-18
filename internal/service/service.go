package service

import (
	"net/http"
	"os"
)

type Server struct {
	Token  string
	ApiKey string
}

func NewServer() *Server {
	return &Server{
		Token:  os.Getenv("SECRET_KEY"),
		ApiKey: os.Getenv("API_KEY"),
	}
}

func Init(s *Server) {
	http.Handle("/", Middleware(http.HandlerFunc(s.initializeChatOrButton)))
	http.Handle("/:sessionUUID/:user", Middleware(http.HandlerFunc(s.Front)))
	http.Handle("/SetAddress", Middleware(http.HandlerFunc(s.SetAddressAndChat)))
}
