package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Id     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

type BookStore struct {
	store map[string]Book
}

func createBookRepository() *BookStore {
	return &BookStore{
		store: map[string]Book{
			"book1": Book{Id: "1", Title: "Harry Potter", Author: "JK Rowling"},
		},
	}
}

func (bs *BookStore) booksHandler(w http.ResponseWriter, r *http.Request) {
	jsonByte, err := json.Marshal(bs.store)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte)(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

func main() {
	bookRepo := createBookRepository()
	http.HandleFunc("/books", bookRepo.booksHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listen on port: 8080")
}
