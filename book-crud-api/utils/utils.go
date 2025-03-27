package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"book-crud-api/models" 
)

// ReadBooks reads book data from the JSON file
func ReadBooks(filename string) ([]models.Book, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var books []models.Book
    bytes, _ := ioutil.ReadAll(file)
    json.Unmarshal(bytes, &books)
    return books, nil
}

// WriteBooks writes the book data to the JSON file
func WriteBooks(filename string, books []models.Book) error {
    bytes, err := json.MarshalIndent(books, "", "  ") // Pretty-print JSON
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filename, bytes, 0644)
}