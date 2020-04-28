package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-musings/server/database"
	"github.com/go-musings/server/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env file not found")
		os.Exit(1)
	}
}

var (
	pgUser     = os.Getenv("PG_USER")
	pgPassword = os.Getenv("PG_PASSWORD")
	pgDatabase = os.Getenv("PG_DATABASE")
)

func main() {
	connString := fmt.Sprintf("postgres://%s:%s@localhost/%s", pgUser, pgPassword, pgDatabase)
	err := database.ConnectDB("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	// fs := http.FileServer(http.Dir("static/")) // file handler
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter() // create mux router

	// routes - using the new mux router
	// r.HandleFunc("/", rootHandler)
	// create book router
	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/", handlers.GetBooksHandler).Methods("GET")
	bookRouter.HandleFunc("/{id}", handlers.GetBookHandler).Methods("GET")
	bookRouter.HandleFunc("/", handlers.CreateBookHandler).Methods("POST")
	bookRouter.HandleFunc("/{id}", handlers.UpdateBookHandler).Methods("PUT")
	bookRouter.HandleFunc("/{id}", handlers.DeleteBookHandler).Methods("DELETE")

	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r)) // pass the mux router to the listener
}
