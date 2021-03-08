package main

import (
	"log"
	"net/http"
    "bookshelf/driver"
    "bookshelf/controllers"

	"github.com/gorilla/mux"
    "github.com/subosito/gotenv"
)

func init() {
    gotenv.Load()
}

func main() {
    db := driver.ConnectDB()
    router := mux.NewRouter()
    controller := controllers.Controller{}

    router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
    router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
    router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
    router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
    router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

    log.Println("server on :5555")
    log.Fatal(http.ListenAndServe(":5555", router))
}

