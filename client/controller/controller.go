package controller

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
)

var serverAddr = "localhost:8081"

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Proxying request: %s %s\n", r.Method, r.URL.String())

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		http.Error(w, "Failed to connect to offshore server: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer conn.Close()

	err = r.Write(conn)
	if err != nil {
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadGateway)
		return
	}

	resp, err := http.ReadResponse(bufio.NewReader(conn), r)
	if err != nil {
		http.Error(w, "Failed to read response: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
