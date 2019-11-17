package models

// Book in store
type Book struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
	// Author *Author `json:"author"`
}
