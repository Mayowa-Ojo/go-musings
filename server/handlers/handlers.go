package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-musings/server/database"
	models "github.com/go-musings/server/database/models"
	"github.com/go-musings/server/templates"
	"github.com/gorilla/mux"
)

// RootHandler -
func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := templates.ParseTemplate("./templates/index.tpl")

	tmpl.Execute(w, nil)
	// fmt.Fprintf(w, "Hello, you made a Request to %s\n", r.URL.Path)
}

// GetBooksHandler -
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	path := r.URL.Path

	books, err := database.QueryRows(db)
	if err != nil {
		log.Fatal(err)
	}

	if contains := strings.Contains(path, "api"); contains {
		b, err := json.Marshal(books)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

		return
	}

	tmpl := templates.ParseTemplate("./templates/books.tpl")

	tmpl.Execute(w, books)
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
	path := r.URL.Path
	b := models.Book{}

	if contains := strings.Contains(path, "api"); contains {

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

		return
	}

	// get form data
	b.Title = r.FormValue("title")
	b.Author = r.FormValue("author")

	err := database.QueryInsertRow(db, b)

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/books/", http.StatusMovedPermanently)
}

// UpdateBookHandler -
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var update models.Book

	path := r.URL.Path
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)

	if contains := strings.Contains(path, "api"); !contains {
		update.Title = r.FormValue("title")
		update.Author = r.FormValue("author")

		err := database.QueryUpdateRow(db, idInt, "books", update)
		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/books/", http.StatusMovedPermanently)

		return
	}

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

	path := r.URL.Path
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)

	err := database.QueryDeleteRow(db, idInt, "books")

	if err != nil {
		log.Fatal(err)
	}

	if contains := strings.Contains(path, "api"); !contains {
		http.Redirect(w, r, "/books/", http.StatusMovedPermanently)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book successfully deleted..."))
}

// ShowFormHandler -
func ShowFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := templates.ParseTemplate("./templates/form.tpl")
	path := r.URL.Path

	if contains := strings.Contains(path, "new"); contains {
		tmpl.Execute(w, nil)
		return
	}

	db := database.GetDB()
	params := mux.Vars(r)
	id := params["id"]
	idInt, _ := strconv.Atoi(id)

	book, err := database.QueryRow(db, "books", idInt)

	if err != nil {
		log.Fatal(err)
	}

	if contains := strings.Contains(path, "edit"); contains {
		action := fmt.Sprintf("/books/%s", id)
		data := struct {
			Method string
			Action string
			Book   models.Book
		}{
			Method: "PUT",
			Action: action,
			Book:   book,
		}

		tmpl.Execute(w, data)
		return
	}

	// tmpl.Execute(w, nil)
}
