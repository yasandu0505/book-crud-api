package handlers

import (
	"book-crud-api/models"
	"book-crud-api/utils"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

// HandleBooks handles CRUD operations on books
func HandleBooks(w http.ResponseWriter, r *http.Request) {

	// getting the file path dynamically
	cwd, _ := os.Getwd()
	var filename = filepath.Join(cwd, "data", "book.json")

	switch r.Method {
	case "GET":
		books, err := utils.ReadBooks(filename)
		if err != nil {
			http.Error(w, "Unable to read books", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	case "POST":
		var book models.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		books, _ := utils.ReadBooks(filename) // Read current books
		books = append(books, book)           // Add the new book
		utils.WriteBooks(filename, books)     // Write updated list to file

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book) // Respond with the created book
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
