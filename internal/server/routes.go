package server

import (
	"github.com/RomaBiliak/lets-go-chat/internal/auth"
	"github.com/RomaBiliak/lets-go-chat/internal/user"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/v1/user", user.CreateUser).Methods("POST")
	s.Router.HandleFunc("/v1/user/login", auth.Login).Methods("POST")

}
