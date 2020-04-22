package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book - book type
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Books - collection of books
type Books struct {
	Data []Book `json:"data"`
}

var books = Books{
	Data: []Book{
		Book{
			ID:     "1",
			Title:  "Angels and Demons",
			Author: "Dan Brown",
		},
		Book{
			ID:     "2",
			Title:  "Inferno",
			Author: "Dan Brown",
		},
		Book{
			ID:     "3",
			Title:  "The Pelican Brief",
			Author: "John Grisham",
		},
		Book{
			ID:     "4",
			Title:  "The Reckoning",
			Author: "John Grisham",
		},
	},
}

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

// http handlers
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you made a Request to %s\n", r.URL.Path)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(books)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	book, _ := fetchBook(id)

	if book.ID == "" {
		fmt.Fprintf(w, "Book with id: '%v' not found\n", id)
		return
	}

	b, err := json.Marshal(book)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	// fmt.Fprintf(w, "Title: %v | Author: %v\n", book.Title, book.Author)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	b := Book{}

	err := json.NewDecoder(r.Body).Decode(&b) // parse json request body

	if err != nil {
		log.Fatal(err)
	}

	book, err := json.Marshal(b)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(book)
	// fmt.Fprintf(w, "book created\n%v\v", book)

}

func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	update := Book{}

	err := json.NewDecoder(r.Body).Decode(&update)

	if err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r)
	id := params["id"]
	// n, _ := strconv.ParseInt(id, 10, 64) // convert id to int

	book, index := fetchBook(id)

	if book.ID == "" {
		fmt.Fprintf(w, "Book with id: '%v' not found\n", id)
		return
	}

	update.ID = id
	books.Data[*index] = update

	b, err := json.Marshal(books)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, index := fetchBook(id)

	if index == nil {
		w.Write([]byte("Book doesn't exist..."))
		return
	}

	booksSlice := books.Data[:]
	booksSlice = append(booksSlice[:*index], booksSlice[*index+1:len(booksSlice)]...)

	b, err := json.Marshal(booksSlice)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func fetchBook(id string) (b Book, index *int) {

	for i, v := range books.Data {
		if v.ID == id {
			b = books.Data[i]
			index = &i
			break
		}
	}

	return b, index
}
