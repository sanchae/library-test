package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sanchae/library-test/models"
)


func books(router chi.Router) {
	router.Get("/", getAllBooks)
	router.Post("/lend", lendBook)
	router.Post("/return", returnBook)
}

func getAllBooks(w http. ResponseWriter, r *http.Request) {
	books, err := dbInstance.GetAllBooks()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, books); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func lendBook(w http.ResponseWriter, r *http.Request) {
	data := &models.LendBookRequest{}
    if err := render.Bind(r, data); err != nil {
        render.Render(w, r, ErrBadRequest)
        return
    }
	err := dbInstance.LendBook(data.BookID, data.UserID)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	render.Status(r, http.StatusOK)
}

func returnBook(w http.ResponseWriter, r *http.Request) {
	data := &models.ReturnBookRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	err := dbInstance.ReturnBook(data.BookID, data.UserID)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	render.Status(r, http.StatusOK)
}