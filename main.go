package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// dbConfig := DbConfiguration{
	// 	Host:     "localhost",
	// 	User:     "postgres",
	// 	Password: "qwerty123",
	// 	Schema:   "golang",
	// }

	dbConfig := DbConfiguration{}
	readConfiguration("Db", &dbConfig)

	handler := NewHandler(&dbConfig)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.homeLink).Methods(http.MethodGet)
	router.HandleFunc("/book", handler.createBook).Methods(http.MethodPost)
	router.HandleFunc("/books", handler.getAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.getBookByID).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.updateExistingBook).Methods(http.MethodPatch)
	router.HandleFunc("/books/{id}", handler.deleteBook).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":18132", router))
}
