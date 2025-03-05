package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	
	log.Print("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}