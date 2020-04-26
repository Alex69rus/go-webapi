package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Alex69rus/go-webapi/models"
	"github.com/Alex69rus/go-webapi/repo"

	"github.com/gorilla/mux"
)

// Handler is http requests handler
type Handler struct {
	repo repo.BookRepository
}

// NewHandler creates new Handler
func NewHandler(bookRepo repo.BookRepository) *Handler {
	return &Handler{
		repo: bookRepo,
	}
}

func (h *Handler) homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newBook)
	id, err := h.repo.InsertBook(newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	dbBooks := h.repo.GetBooks()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbBooks)
}

func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	id := h.getIntFromStr(mux.Vars(r)["id"])

	book := h.repo.GetBook(id)
	if book != nil {
		json.NewEncoder(w).Encode(book)
		return
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func (h *Handler) updateExistingBook(w http.ResponseWriter, r *http.Request) {
	id := h.getIntFromStr(mux.Vars(r)["id"])

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Kindly enter data with the event title and description only in order to update")
	}
	newBook := models.Book{}
	unmarshalErr := json.Unmarshal(reqBody, &newBook)
	if unmarshalErr != nil {
		fmt.Fprint(w, "unmarshalling error: ", unmarshalErr)
	}
	h.repo.UpdateBook(id, newBook.Title)

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := h.getIntFromStr(mux.Vars(r)["id"])

	h.repo.DeleteBook(id)
}

func (h *Handler) getIntFromStr(str string) int32 {
	eventIDInt64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1
	}
	eventID := int32(eventIDInt64)
	return eventID
}
