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
	case "PUT":
		var updatedBook models.Book
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Extract book ID from URL path (assuming book ID is part of the URL)
		bookID := r.URL.Query().Get("id")
		if bookID == "" {
			http.Error(w, "Book ID is required", http.StatusBadRequest)
			return
		}

		books, err := utils.ReadBooks(filename)
		if err != nil {
			http.Error(w, "Unable to read books", http.StatusInternalServerError)
			return
		}

		// Find and update the book with the matching ID
		var bookFound bool
		for i, book := range books {
			if book.BookID == bookID {
				books[i] = updatedBook
				bookFound = true
				break
			}
		}

		if !bookFound {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		// Write the updated book list back to the file
		err = utils.WriteBooks(filename, books)
		if err != nil {
			http.Error(w, "Unable to write books", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedBook)

	case "DELETE":
		// Extract book ID from URL path (assuming book ID is part of the URL)
		bookID := r.URL.Query().Get("id")
		if bookID == "" {
			http.Error(w, "Book ID is required", http.StatusBadRequest)
			return
		}

		books, err := utils.ReadBooks(filename)
		if err != nil {
			http.Error(w, "Unable to read books", http.StatusInternalServerError)
			return
		}

		// Find and delete the book with the matching ID
		var bookFound bool
		var updatedBooks []models.Book
		for _, book := range books {
			if book.BookID != bookID {
				updatedBooks = append(updatedBooks, book)
			} else {
				bookFound = true
			}
		}

		if !bookFound {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		// Write the updated book list back to the file
		err = utils.WriteBooks(filename, updatedBooks)
		if err != nil {
			http.Error(w, "Unable to write books", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
