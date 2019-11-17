package repo

import (
	"fmt"

	. "github.com/alex69rus/Webapi/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // because sqlx require this package
)

const (
	pgURL = "postgresql://localhost/postgres?user=postgres&password=qwerty123&sslmode=disable&search_path=golang"
)

// GetBooks returns all books
func GetBooks() *[]Book {
	db, err := sqlx.Connect("postgres", pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	dbBooks := []Book{}
	err = db.Select(&dbBooks, "SELECT * FROM books")
	if err != nil {
		fmt.Println("While SELECT occured error: ", err)
	}

	return &dbBooks
}

// GetBook find book by id in DB
func GetBook(id int32) *Book {
	db, err := sqlx.Connect("postgres", pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	dbBook := Book{}
	err = db.Get(&dbBook, "SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		fmt.Println("While SELECT by id occured error: ", err)
		return nil
	}

	return &dbBook
}

// InsertBook insert new book in db
func InsertBook(newTitle string) int32 {
	db, err := sqlx.Connect("postgres", pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	res := db.MustExec("INSERT INTO books (title) VALUES($1)", newTitle)
	newID, _ := res.LastInsertId()
	return int32(newID)
}

// UpdateBook update book in db
func UpdateBook(id int32, newTitle string) {
	db, err := sqlx.Connect("postgres", pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	db.MustExec("UPDATE books SET title=$1 WHERE id=$2", newTitle, id)
}
