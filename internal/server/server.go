package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) Run(addr string) {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
	fmt.Println("Listening to port "+addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}