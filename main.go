package main

import (
    "log"
    "net/http"

    "book-crud-api/handlers"
)

func main() {
    http.HandleFunc("/books", handlers.HandleBooks)
    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
