package controllers

import (
	"log"
    "strconv"
	"net/http"
	"database/sql"
	"encoding/json"
    "bookshelf/models"
    "bookshelf/repository"

	"github.com/gorilla/mux"
)

type Controller struct{}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var book models.Book
        params := mux.Vars(r)

        bookRepo := bookRepository.BookRepository{}

        id, err := strconv.Atoi(params["id"])
        if err != nil {
            log.Fatal(err)
        }

        book = bookRepo.GetBook(db, book, id)
        json.NewEncoder(w).Encode(book)
    }
}


func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var book models.Book
        var books = []models.Book{}

        bookRepo := bookRepository.BookRepository{}
        books = bookRepo.GetBooks(db, book, books)

        json.NewEncoder(w).Encode(books)
    }
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var book models.Book
        var bookID int

        json.NewDecoder(r.Body).Decode(&book)

        bookRepo := bookRepository.BookRepository{}
        bookID = bookRepo.AddBook(db, book)

        json.NewEncoder(w).Encode(bookID)
    }
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var book models.Book
        json.NewDecoder(r.Body).Decode(&book)

        bookRepo := bookRepository.BookRepository{}
        rowsUpdated := bookRepo.UpdateBook(db, book)

        json.NewEncoder(w).Encode(rowsUpdated)
    }
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)

        bookRepo := bookRepository.BookRepository{}

        id, err := strconv.Atoi(params["id"])
        if err != nil {
            log.Fatal(err)
        }

        rowsDeleted := bookRepo.RemoveBook(db, id)

        json.NewEncoder(w).Encode(rowsDeleted)
    }
}
