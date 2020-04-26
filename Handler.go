package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Alex69rus/webapi/models"
	"github.com/Alex69rus/webapi/repo"

	"github.com/gorilla/mux"
)

// Handler is http requests handler
type Handler struct {
	dbConfig *DbConfiguration
}

// NewHandler creates new Handler
func NewHandler(dbConfig *DbConfiguration) *Handler {
	return &Handler{
		dbConfig: dbConfig,
	}
}

func (h *Handler) homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	repo := h.createBookRepository()

	var newBook models.Book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &newBook)
	id, err := repo.InsertBook(newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	repo := h.createBookRepository()

	dbBooks := repo.GetBooks()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbBooks)
}

func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) {
	repo := h.createBookRepository()
	id := h.getIntFromStr(mux.Vars(r)["id"])

	book := repo.GetBook(id)
	if book != nil {
		json.NewEncoder(w).Encode(book)
		return
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func (h *Handler) updateExistingBook(w http.ResponseWriter, r *http.Request) {
	repo := h.createBookRepository()
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
	repo.UpdateBook(id, newBook.Title)

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	repo := h.createBookRepository()
	id := h.getIntFromStr(mux.Vars(r)["id"])

	repo.DeleteBook(id)
}

func (h *Handler) getIntFromStr(str string) int32 {
	eventIDInt64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1
	}
	eventID := int32(eventIDInt64)
	return eventID
}

func (h *Handler) createBookRepository() repo.BookRepository {
	return repo.NewBookRepository(h.dbConfig.Host, h.dbConfig.User, h.dbConfig.Password, h.dbConfig.Schema)
}
