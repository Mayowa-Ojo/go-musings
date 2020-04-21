package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book - map of book type with id as key
type Book map[string]string

// Books - map of Books
type Books map[string]Book

var books = Books{
	"1": Book{
		"title":  "Angels and Demons",
		"author": "Dan Brown",
	},
	"2": Book{
		"title":  "Inferno",
		"author": "Dan Brown",
	},
	"3": Book{
		"title":  "The Pelican Brief",
		"author": "John Grisham",
	},
	"4": Book{
		"title":  "The Reckoning",
		"author": "John Grisham",
	},
}

func main() {
	// fs := http.FileServer(http.Dir("static/")) // file handler
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter() // create mux router

	// using the new mux router
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/books", booksHandler)
	r.HandleFunc("/books/{id}", bookHandler)

	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r)) // pass the mux router to the listener
}

// http handlers
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you made a Request to %s\n", r.URL.Path)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range books {
		fmt.Fprintf(w, "book-id: %v | title: %v | author: %v\n", k, v["title"], v["author"])
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	book, ok := books[id]

	if !ok {
		fmt.Fprintf(w, "Book with id: '%v' not found\n", id)
		return
	}
	fmt.Fprintf(w, "Title: %v | Author: %v\n", book["title"], book["author"])
}
