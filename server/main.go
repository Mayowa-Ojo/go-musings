package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-musings/server/database"
	"github.com/go-musings/server/handlers"
	"github.com/go-musings/server/middleware"
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

func main() {
	var (
		pgUser     = os.Getenv("PG_USER")
		pgPassword = os.Getenv("PG_PASSWORD")
		pgDatabase = os.Getenv("PG_DATABASE")
		port       = os.Getenv("PORT")
	)
	connString := fmt.Sprintf("postgres://%s:%s@localhost/%s", pgUser, pgPassword, pgDatabase)
	err := database.ConnectDB("postgres", connString)
	addr := fmt.Sprintf(":%s", port)

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter() // create mux router

	fs := http.FileServer(http.Dir("./static")) // file handler
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// routes - using the new mux router
	r.HandleFunc("/", handlers.RootHandler)

	bookRouter := r.PathPrefix("/books").Subrouter() // create book router
	apiRouter := r.PathPrefix("/api/v1").Subrouter() // create api router

	bookRouter.HandleFunc("/", handlers.GetBooksHandler).Methods("GET")
	bookRouter.HandleFunc("/edit/{id}", handlers.ShowFormHandler).Methods("GET")
	bookRouter.HandleFunc("/new", handlers.ShowFormHandler).Methods("GET")
	bookRouter.HandleFunc("/{id}", handlers.GetBookHandler).Methods("GET")
	bookRouter.HandleFunc("/", handlers.CreateBookHandler).Methods("POST")
	bookRouter.HandleFunc("/{id}", handlers.UpdateBookHandler).Methods("PUT")
	bookRouter.HandleFunc("/delete/{id}", handlers.DeleteBookHandler).Methods("GET") // using GET here because the CTA is an anchor tag

	apiRouter.HandleFunc("/books", handlers.GetBooksHandler).Methods("GET")
	apiRouter.HandleFunc("/books/{id}", handlers.GetBookHandler).Methods("GET")
	apiRouter.HandleFunc("/{id}", handlers.DeleteBookHandler).Methods("DELETE")

	// http.Handle("/", r)
	log.Printf("http server listening on port %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, middleware.MethodOverride(r))) // pass the mux router to the listener
}
