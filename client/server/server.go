package server

import (
	"log"
	"net/http"
	"ship-proxy/controller"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	r := chi.NewRouter()

	r.HandleFunc("/*", controller.ProxyHandler)

	return &Server{router: r}

}

func (s *Server) Start() {
	log.Println("SHIP proxy running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
