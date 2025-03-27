package handlers

import (
	"encoding/json" // Keep it, now we’re using it!
	"fmt"
	"net/http"
)

// Temporary sample structure (can be adjusted later)
type Book struct {
    Title string `json:"title"`
}

// HandleBooks handles the /books endpoint
func HandleBooks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        books := []Book{
            {Title: "The Great Gatsby"},
            {Title: "To Kill a Mockingbird"},
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(books) // Write JSON response
    case "POST":
        var book Book
        err := json.NewDecoder(r.Body).Decode(&book)
        if err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        fmt.Fprintf(w, "Book received: %+v\n", book) // Respond with received data
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
