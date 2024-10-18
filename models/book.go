package models

import (
	"net/http"
)

type Book struct {
	ID int `json:"id"`
	BookTitle string `json:"book_title"`
	AvailableCopies string `json:"available_copies"`
}

type BookList struct {
	Books []Book `json:"books"`
}

type LendBookRequest struct {
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
}

type ReturnBookRequest struct {
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
}

func (*BookList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Book) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (l *LendBookRequest) Bind(r *http.Request) error {
	return nil
}

func (l *ReturnBookRequest) Bind(r *http.Request) error {
	return nil
}