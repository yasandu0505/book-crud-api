package handlers

import (
	"book-crud-api/models"
	"book-crud-api/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const filename = "data/book.json" // JSON file location

// HandleBooks - Handles CRUD operations on books
func HandleBooks(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/books/")
	if path == "" { // If path is empty, handle `/books` CRUD operations
		switch r.Method {
		case "GET":
			getAllBooks(w)
		case "POST":
			createBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else { // Handle `/books/{id}` for GET, PUT, and DELETE
		bookID := path
		switch r.Method {
		case "GET":
			getBookByID(w, bookID)
		case "PUT":
			updateBookByID(w, r, bookID)
		case "DELETE":
			deleteBookByID(w, bookID)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// Helper functions for CRUD operations

// getAllBooks - Handle GET /books
func getAllBooks(w http.ResponseWriter) {
	books, err := utils.ReadBooks(filename)
	if err != nil {
		http.Error(w, "Unable to read books", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

// createBook - Handle POST /books
func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	json.NewDecoder(r.Body).Decode(&newBook)

	books, err := utils.ReadBooks(filename)
	if err != nil {
		http.Error(w, "Unable to read books", http.StatusInternalServerError)
		return
	}

	books = append(books, newBook)
	utils.WriteBooks(filename, books)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

// getBookByID - Handle GET /books/{id}
func getBookByID(w http.ResponseWriter, bookID string) {
	books, err := utils.ReadBooks(filename)
	if err != nil {
		http.Error(w, "Unable to read books", http.StatusInternalServerError)
		return
	}

	for _, book := range books {
		if book.BookID == bookID {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// updateBookByID - Handle PUT /books/{id}
func updateBookByID(w http.ResponseWriter, r *http.Request, bookID string) {
	books, err := utils.ReadBooks(filename)
	if err != nil {
		http.Error(w, "Unable to read books", http.StatusInternalServerError)
		return
	}

	var updatedBook models.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	for i, book := range books {
		if book.BookID == bookID {
			books[i] = updatedBook // Update the book details in the slice
			utils.WriteBooks(filename, books)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

// deleteBookByID - Handle DELETE /books/{id}
func deleteBookByID(w http.ResponseWriter, bookID string) {
	books, err := utils.ReadBooks(filename)
	if err != nil {
		http.Error(w, "Unable to read books", http.StatusInternalServerError)
		return
	}

	for i, book := range books {
		if book.BookID == bookID {
			books = append(books[:i], books[i+1:]...) // Remove book from slice
			utils.WriteBooks(filename, books)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Book deleted successfully")
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
