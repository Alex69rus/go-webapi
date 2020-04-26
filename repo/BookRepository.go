package repo

import (
	"github.com/Alex69rus/go-webapi/models"
)

// BookRepository provide access to books
type BookRepository interface {
	GetBooks() *[]models.Book
	GetBook(id int32) *models.Book
	InsertBook(newBook models.Book) (int32, error)
	UpdateBook(id int32, newTitle string)
	DeleteBook(id int32)
}
