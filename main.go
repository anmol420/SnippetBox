package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	// NOT FOUND PAGE
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet View"))
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

func main() {
	mux := http.NewServeMux() // create a router
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		println("Error starting server")
	}
}