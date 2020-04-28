package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-musings/server/database"
	models "github.com/go-musings/server/database/models"
	"github.com/gorilla/mux"
)

// RootHandler -
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you made a Request to %s\n", r.URL.Path)
}

// GetBooksHandler -
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	books, err := database.QueryRows(db)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// GetBookHandler -
func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)

	book, err := database.QueryRow(db, "books", idInt)

	if err != nil {
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
}

// CreateBookHandler -
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	b := models.Book{}

	err := json.NewDecoder(r.Body).Decode(&b) // parse json request body
	if err != nil {
		log.Fatal(err)
	}

	err = database.QueryInsertRow(db, b)
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
}

// UpdateBookHandler -
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	update := models.Book{}
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Fatal(err)
	}

	err = database.QueryUpdateRow(db, idInt, "books", update)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(update)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// DeleteBookHandler -
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)

	err := database.QueryDeleteRow(db, idInt, "books")

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book successfully deleted..."))
}
