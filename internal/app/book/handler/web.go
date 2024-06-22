package handler

import (
	"github.com/go-chi/chi/v5"
	"rest-api/internal/app/book/usecase"
)

type BookRegistryOpts struct {
	Book usecase.Book
}

type BookRegistry struct {
	opts BookRegistryOpts
}

func NewBookRegistry(b BookRegistryOpts) *BookRegistry {
	return &BookRegistry{
		b,
	}
}

func (s BookRegistry) RegisterRoutesTo(r chi.Router) error {
	r.Get("/books", GetBooks(s.opts.Book))
	r.Get("/book/{id}", GetBookDetail(s.opts.Book))
	return nil
}
