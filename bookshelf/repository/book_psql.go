package bookRepository

import (
    "log"
    "database/sql"
    "bookshelf/models"
)

type BookRepository struct{}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
    rows := db.QueryRow("SELECT * FROM books WHERE id=$1", id)

    err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
    if err != nil {
        log.Fatal(err)
    }

    return book
}


func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
        if err != nil {
            log.Fatal(err)
        }
        books = append(books, book)
    }

    return books
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) int {
    err := db.QueryRow("INSERT INTO books (title, author, year) VALUES($1, $2, $3) RETURNING id", book.Title, book.Author, book.Year).Scan(&book.ID)

    if err != nil {
        log.Fatal(err)
    }

    return book.ID
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
    result, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)
    if err != nil {
        log.Fatal(err)
    }

    rowsUpdated, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }

    return rowsUpdated
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
    result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
    if err != nil {
        log.Fatal(err)
    }

    rowsDeleted, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }

    return rowsDeleted
}
