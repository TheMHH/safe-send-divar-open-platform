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

	mediaS := service.NewServer()

	service.Init(mediaS)

	if err := http.Serve(lis, nil); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func main() {
	StartServer()
}
