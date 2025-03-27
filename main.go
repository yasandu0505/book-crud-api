package main

import (
	"book-crud-api/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/books/search", handlers.HandleBooksSearch)
	http.HandleFunc("/books", handlers.HandleBooks)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
