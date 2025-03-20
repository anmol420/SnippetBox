package main

import (
	"log"
	"net/http"
	"os"

	"github.com/anmol420/LearnGoBackend/internal/env"
)

// application struct for dependency injection
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	err := env.LoadEnv()
	if err != nil {
		errorLogger.Fatal("Error in loading .env")
	}

	app := &application{
		errorLog: errorLogger,
		infoLog: infoLogger,
	}

	// custom http.Server struct
	srv := &http.Server{
		Addr: os.Getenv("ADDR"),
		ErrorLog: errorLogger,
		Handler: app.routes(),
	}

	infoLogger.Printf("Server is running on port %s", os.Getenv("ADDR"))
	listenErr := srv.ListenAndServe() // calling custom struct instead of http package
	if listenErr != nil {
		errorLogger.Fatal(listenErr)
	}
}