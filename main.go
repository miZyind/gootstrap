package main

import (
	"log"
	"net/http"
)

func main() {
	// Setup web server
	server := &http.Server{
		Addr: ":8080",
	}
	// Setup file server
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Serve
	log.Printf("Gootstrap server is started at %v", server.Addr)
	panic(server.ListenAndServe())
}
