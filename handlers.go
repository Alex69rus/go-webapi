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

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newBook)
	id, err := InsertBook(newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	dbBooks := GetBooks()
	json.NewEncoder(w).Encode(dbBooks)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	id := getIntFromStr(mux.Vars(r)["id"])

	book := GetBook(id)
	if book != nil {
		json.NewEncoder(w).Encode(book)
		return
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
	id := getIntFromStr(mux.Vars(r)["id"])

	DeleteBook(id)
}
func getIntFromStr(str string) int32 {
	eventIDInt64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1
	}
	eventID := int32(eventIDInt64)
	return eventID
}
