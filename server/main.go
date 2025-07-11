package main

import (
	"log"
	"net"
	"server/service"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Failed to start offshore proxy server:", err)
	}
	log.Println("Offshore proxy server listening on :8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go service.HandleConnection(conn)
	}
}
