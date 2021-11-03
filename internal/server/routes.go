package server

import (
	"github.com/RomaBiliak/lets-go-chat/internal/item"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/items", item.Items).Methods("GET")
	s.Router.HandleFunc("/item/{id}", item.Item).Methods("GET")
	s.Router.HandleFunc("/item", item.CreateItem).Methods("POST")
	s.Router.HandleFunc("/item/{id}", item.UpdateItem).Methods("PUT")
	s.Router.HandleFunc("/item/{id}", item.DeleteItem).Methods("DELETE")

}
