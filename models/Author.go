package models

// Author is author of book
type Author struct {
	Firstname string `json:"firstname" db:"first_name"`
	Lastname  string `json:"lastname" db:"last_name"`
}
