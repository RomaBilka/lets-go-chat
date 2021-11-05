package http

import (
	"fmt"
	"log"
	"net/http"
)

func Start (addr string) {
	s:= &http.Server{
		Addr: addr,
	}
	fmt.Println("Listening "+addr)
	log.Fatal(s.ListenAndServe())
}
