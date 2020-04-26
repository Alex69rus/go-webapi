package repo

import (
	"fmt"

	. "github.com/Alex69rus/webapi/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // because sqlx require this package
)

// BookRepositoryImpl provide access to books
type BookRepositoryImpl struct {
	pgURL string
}

// NewBookRepository creates new BookRepository
func NewBookRepository(host string, user string, password string, schema string) *BookRepositoryImpl {
	connectStr := fmt.Sprintf("postgresql://%v/postgres?user=%v&password=%v&sslmode=disable&search_path=%v",
		host, user, password, schema)
	return &BookRepositoryImpl{
		pgURL: connectStr,
	}
}

const (
	selectSQL      = "SELECT * FROM books"
	selectWhereSQL = "SELECT * FROM books WHERE id=$1"
	insertSQL      = "INSERT INTO books (title) VALUES(:title) RETURNING id"
	updateSQL      = "UPDATE books SET title=$1 WHERE id=$2"
	deleteSQL      = "DELETE FROM books WHERE id=$1"
)

// GetBooks returns all books
func (repo *BookRepositoryImpl) GetBooks() *[]Book {
	db, err := sqlx.Connect("postgres", repo.pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	dbBooks := []Book{}
	err = db.Select(&dbBooks, selectSQL)
	if err != nil {
		fmt.Println("While SELECT occured error: ", err)
	}

	return &dbBooks
}

// GetBook find book by id in DB
func (repo *BookRepositoryImpl) GetBook(id int32) *Book {
	db, err := sqlx.Connect("postgres", repo.pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	dbBook := Book{}
	err = db.Get(&dbBook, selectWhereSQL, id)
	if err != nil {
		fmt.Println("While SELECT by id occured error: ", err)
		return nil
	}

	return &dbBook
}

// InsertBook insert new book in db
func (repo *BookRepositoryImpl) InsertBook(newBook Book) (int32, error) {
	db, err := sqlx.Connect("postgres", repo.pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	rows, err := db.NamedQuery(insertSQL, newBook)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
		return 0, fmt.Errorf("Coudn't create book. Reason:%v", err)
	}

	var newID int
	if rows.Next() {
		rows.Scan(&newID)
	}

	return int32(newID), nil
}

// UpdateBook update book in db
func (repo *BookRepositoryImpl) UpdateBook(id int32, newTitle string) {
	db, err := sqlx.Connect("postgres", repo.pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	db.MustExec(updateSQL, newTitle, id)
}

// DeleteBook delete book in db by id
func (repo *BookRepositoryImpl) DeleteBook(id int32) {
	db, err := sqlx.Connect("postgres", repo.pgURL)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	db.MustExec(deleteSQL, id)
}
