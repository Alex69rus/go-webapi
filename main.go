package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink).Methods(http.MethodGet)
	router.HandleFunc("/book", createBook).Methods(http.MethodPost)
	router.HandleFunc("/books", getAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", getBookByID).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", updateExistingBook).Methods(http.MethodPatch)
	router.HandleFunc("/books/{id}", deleteBook).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":18132", router))
}
