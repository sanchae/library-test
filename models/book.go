package models

import (
	"fmt"
	"net/http"
)

type Book struct {
	ID int `json:"id"`
	BookTitle string `json:"book_title"`
	AvailableCopies string `json:"available_copies"`
}

type BookList struct {
	BookS []Book `json:"books"`
}

func (i *Book) Bind(r *http.Request) error {
	if i.BookTitle == "" {
		return fmt.Errorf("booktitle is a required field")
	}
	if i.AvailableCopies == "" {
		return fmt.Errorf("available copies is a required field")
	}
	return nil
}
func (*BookList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Book) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}