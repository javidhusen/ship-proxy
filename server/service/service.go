package service

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Println("Failed to read request:", err)
		return
	}

	if req.URL.Scheme == "" || req.URL.Host == "" {
		req.URL.Scheme = "http"
		req.URL.Host = req.Host
	}

	req.RequestURI = ""

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to forward request:", err)
		return
	}
	defer resp.Body.Close()

	err = resp.Write(conn)
	if err != nil {
		log.Println("Failed to write response:", err)
	}
}
