package http

import (
	"fmt"
	"log"
	"net/http"
)

func Start(addr string, mux *http.ServeMux) {
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	fmt.Printf("Listening %s\n", addr)
	log.Fatal(s.ListenAndServe())
}
