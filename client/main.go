package main

import (
	"ship-proxy/server"
)

func main() {

	srv := server.NewServer()

	srv.Start()

}
