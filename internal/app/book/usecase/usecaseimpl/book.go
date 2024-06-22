package usecaseimpl

import "context"

type Book struct{}

func NewBook(b Book) Book {
	return b
}

func (b Book) Get(ctx context.Context) []string {
	books := []string{
		"newspaper",
		"tabloid",
	}
	return books
}

func (b Book) GetByID(ctx context.Context, id string) string {
	switch id {
	case "1":
		return "newspaper"
	case "2":
		return "tabloid"
	default:
		return "not found"
	}
}
