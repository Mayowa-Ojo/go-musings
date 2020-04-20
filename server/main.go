package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fs := http.FileServer(http.Dir("static/")) // file handler
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter() // create mux router

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// http handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you made a Request to %s\n", r.URL.Path)
}
