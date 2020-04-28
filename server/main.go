package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// fs := http.FileServer(http.Dir("static/")) // file handler
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter() // create mux router

	// routes - using the new mux router
	r.HandleFunc("/", rootHandler)
	// create book router
	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/", getBooksHandler).Methods("GET")
	bookRouter.HandleFunc("/{id}", getBookHandler).Methods("GET")
	bookRouter.HandleFunc("/", createBookHandler).Methods("POST")
	bookRouter.HandleFunc("/{id}", updateBookHandler).Methods("PUT")
	bookRouter.HandleFunc("/{id}", deleteBookHandler).Methods("DELETE")

	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r)) // pass the mux router to the listener
}
