package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// NOT FOUND PAGE
	if r.URL.Path != "/" {
		app.notFound(w) // custom errors
		return
	}

	files := []string{
		"./ui/html/base.tmpl", // it is important that first file must be the base file
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	// template parsing
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Print(err.Error())
		app.serverError(w, err)
		return
	}
	// execute the template set with a name -> "base"
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Print(err.Error())
		app.serverError(w, err)
	}
	// w.Write([]byte("Hello World"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	// w.Write([]byte("Snippet View"))
	fmt.Fprintf(w, "ID: %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405) // once per response
		// w.Write([]byte("Must Be POST Method"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Snippet Create"))
}