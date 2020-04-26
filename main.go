package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Alex69rus/go-webapi/repo"
	"github.com/gorilla/mux"
)

var dbConfig *repo.DbConfiguration = &repo.DbConfiguration{}

func init() {
	log.SetOutput(os.Stdout)
	readConfiguration("Db", dbConfig)
}

func main() {
	bookRepo := repo.NewBookRepository(dbConfig)
	handler := NewHandler(bookRepo)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.homeLink).Methods(http.MethodGet)
	router.HandleFunc("/book", handler.createBook).Methods(http.MethodPost)
	router.HandleFunc("/books", handler.getAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.getBookByID).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handler.updateExistingBook).Methods(http.MethodPatch)
	router.HandleFunc("/books/{id}", handler.deleteBook).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":18132", router))
}
