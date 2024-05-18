package main

import (
	"log"
	"main/internal/service"
	"net"
	"net/http"
)

func StartServer() {
	lis, err := net.Listen("tcp", ":9290")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		return
	}

	safeSendS := service.NewServer()

	service.Init(safeSendS)

	if err := http.Serve(lis, nil); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func main() {
	StartServer()
}
