package main

import (
	"log"
	"net/http"

	"github.com/go-musings/server/database"
	"github.com/go-musings/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	err := database.ConnectDB("postgres", "postgres://mayor:adebayor@localhost/test")

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
