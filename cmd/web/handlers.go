package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// NOT FOUND PAGE
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// template parsing
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// execute the template set
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// w.Write([]byte("Hello World"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Snippet View"))
	fmt.Fprintf(w, "ID: %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405) // once per response
		// w.Write([]byte("Must Be POST Method"))
		http.Error(w, "Must Be POST Method", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Snippet Create"))
}