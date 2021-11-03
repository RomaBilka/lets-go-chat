package server

import (
	"flag"
	"github.com/RomaBiliak/lets-go-chat/internal/item"
	"net/http"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/items", item.Items).Methods("GET")
	s.Router.HandleFunc("/item/{id}", item.Item).Methods("GET")
	s.Router.HandleFunc("/item", item.CreateItem).Methods("POST")
	s.Router.HandleFunc("/item/{id}", item.UpdateItem).Methods("PUT")
	s.Router.HandleFunc("/item/{id}", item.DeleteItem).Methods("DELETE")

	var dir string
	flag.StringVar(&dir, "api", "api/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	s.Router.PathPrefix("/api/").Handler(http.StripPrefix("/api/", http.FileServer(http.Dir(dir))))

}
