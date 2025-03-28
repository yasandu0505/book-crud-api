package handlers

import (
	"book-crud-api/models"
	"book-crud-api/utils"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// HandleBooksSearch handles the search endpoint: GET /books/search?q=<keyword>
func HandleBooksSearch(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	var filename = filepath.Join(cwd, "data", "book.json")

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	books, err := utils.ReadBooks(filename)
	if err != nil {
		http.Error(w, "Unable to read books", http.StatusInternalServerError)
		return
	}

	// Create a channel to collect matching books from each Goroutine
	resultsChan := make(chan []models.Book)
	var wg sync.WaitGroup

	// Divide books into smaller batches for concurrent search
	batchSize := len(books) / 4 // Split into 4 batches (can be tuned)
	if batchSize == 0 {
		batchSize = 1 // Ensure batchSize is at least 1 for small datasets
	}

	for i := 0; i < len(books); i += batchSize {
		end := i + batchSize
		if end > len(books) {
			end = len(books)
		}

		// Launch a Goroutine to search within each batch
		wg.Add(1)
		go func(batch []models.Book) {
			defer wg.Done()

			var batchMatches []models.Book
			for _, book := range batch {
				if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) ||
					strings.Contains(strings.ToLower(book.Description), strings.ToLower(query)) {
					batchMatches = append(batchMatches, book)
				}
			}
			resultsChan <- batchMatches // Send the matching books to the channel
		}(books[i:end])
	}

	// Goroutine to close the channel after all searches are complete
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect and merge all results from the channel
	var finalResults []models.Book
	for batchResults := range resultsChan {
		finalResults = append(finalResults, batchResults...)
	}

	// Respond with the final aggregated results
	if len(finalResults) == 0 {
		http.Error(w, "No books found matching the query", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(finalResults)
}

// HandleBooks handles CRUD operations on books
func HandleBooks(w http.ResponseWriter, r *http.Request) {

	// getting the file path dynamically
	// cwd, _ := os.Getwd()
	// var filename = filepath.Join(cwd, "data", "book.json")

	//  file path to fetch the data
	var filename = "../data/book.json"

	// file path to test the GET req
	// var filename = "../data/book_test.json"

	switch r.Method {
	case "GET":
		bookID := r.URL.Query().Get("id") // Check if `id` is passed in the query parameters

		books, err := utils.ReadBooks(filename)
		if err != nil {
			http.Error(w, "Unable to read books", http.StatusInternalServerError)
			return
		}

		if bookID != "" {
			// If `id` is provided, find the specific book
			for _, book := range books {
				if book.BookID == bookID {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(book) // Return the specific book
					return
				}
			}
			// If book is not found, return 404
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			// If `id` is not provided, return the entire list
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(books)
		}
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

		var bookFound bool
		for i, book := range books {
			if book.BookID == bookID {
				// Merge the updatedBook fields dynamically with the existing book
				if updatedBook.Title != "" {
					books[i].Title = updatedBook.Title
				}
				if updatedBook.AuthorID != "" {
					books[i].AuthorID = updatedBook.AuthorID
				}
				if updatedBook.PublisherID != "" {
					books[i].PublisherID = updatedBook.PublisherID
				}
				if updatedBook.PublicationDate != "" {
					books[i].PublicationDate = updatedBook.PublicationDate
				}
				if updatedBook.ISBN != "" {
					books[i].ISBN = updatedBook.ISBN
				}
				if updatedBook.Pages != 0 {
					books[i].Pages = updatedBook.Pages
				}
				if updatedBook.Genre != "" {
					books[i].Genre = updatedBook.Genre
				}
				if updatedBook.Description != "" {
					books[i].Description = updatedBook.Description
				}
				if updatedBook.Price != 0 {
					books[i].Price = updatedBook.Price
				}
				if updatedBook.Quantity != 0 {
					books[i].Quantity = updatedBook.Quantity
				}
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
		json.NewEncoder(w).Encode(books) // Respond with the updated book list

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
