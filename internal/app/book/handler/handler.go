package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"rest-api/internal/app/book/usecase"
)

func GetBooks(book usecase.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		list := book.Get(ctx)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, list)
	}
}

func GetBookDetail(book usecase.Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "id")
		list := book.GetByID(ctx, id)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, list)
	}
}
