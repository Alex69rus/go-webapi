package repo

import (
	. "github.com/Alex69rus/webapi/models"
)

// BookRepository provide access to books
type BookRepository interface {
	GetBooks() *[]Book
	GetBook(id int32) *Book
	InsertBook(newBook Book) (int32, error)
	UpdateBook(id int32, newTitle string)
	DeleteBook(id int32)
}
