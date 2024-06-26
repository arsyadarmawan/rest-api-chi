package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var fakeBooks = []*Book{{
	ID:               "1",
	Title:            "7 Habits of Highly Effective People",
	Author:           "Stephen Covey",
	PublishedDate:    "15/08/1989",
	OriginalLanguage: "English",
}}

type fakeStorage struct {
}

func (s fakeStorage) Get(_ string) *Book {
	return fakeBooks[0]
}

func (s fakeStorage) Delete(_ string) *Book {
	return nil
}

func (s fakeStorage) List() []*Book {
	return fakeBooks
}

func (s fakeStorage) Create(_ Book) {
	return
}

func (s fakeStorage) Update(string, Book) *Book {
	return fakeBooks[1]
}

func TestGetBooksHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	w := httptest.NewRecorder()
	bookHandler := BookHandler{
		storage: fakeStorage{},
		//storage: BookStore{},
	}
	bookHandler.GetBooks(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	book := Book{}
	json.Unmarshal(data, &book)
	if book.Title != "7 Habits of Highly Effective People" {
		t.Errorf("expected ABC got %v", string(data))
	}
}
