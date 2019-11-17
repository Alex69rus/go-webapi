package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	. "github.com/alex69rus/Webapi/models"
	. "github.com/alex69rus/Webapi/repo"
	"github.com/gorilla/mux"
)

var books = []Book{
	{ID: 1, Title: "Title1"},
	{ID: 2, Title: "Title2"},
	{ID: 3, Title: "Title3"},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home! Book: %v", books[0])
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newBook)
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	headers := w.Header()
	headers.Add("Content-Type", "application/json")
	dbBooks := GetBooks()
	json.NewEncoder(w).Encode(dbBooks)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	id := getIntFromStr(mux.Vars(r)["id"])

	book := GetBook(id)
	if book != nil {
		json.NewEncoder(w).Encode(book)
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func updateExistingBook(w http.ResponseWriter, r *http.Request) {
	id := getIntFromStr(mux.Vars(r)["id"])

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Kindly enter data with the event title and description only in order to update")
	}
	newBook := Book{}
	unmarshalErr := json.Unmarshal(reqBody, &newBook)
	if unmarshalErr != nil {
		fmt.Fprint(w, "unmarshalling error: ", unmarshalErr)
	}
	UpdateBook(id, newBook.Title)

	w.WriteHeader(http.StatusOK)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	eventIDStr := mux.Vars(r)["id"]
	eventIDInt64, _ := strconv.ParseInt(eventIDStr, 10, 32)
	eventID := int32(eventIDInt64)

	for i, book := range books {
		if book.ID == eventID {
			books = append(books[:i], books[i+1:]...)
			fmt.Fprintf(w, "The book with ID %v has been deleted", eventID)
		}
	}
}

func findBookByID(id string) (book *Book, err error) {
	eventIDInt64, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Could not parse id from query. Error: %v", err)
	}
	eventID := int32(eventIDInt64)

	for i := range books {
		if books[i].ID == eventID {
			return &books[i], nil
		}
	}
	return nil, nil
}

func getIntFromStr(str string) int32 {
	eventIDInt64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1
	}
	eventID := int32(eventIDInt64)
	return eventID
}
