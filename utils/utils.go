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
    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(bytes, &books)
    if err != nil {
        return nil, err
    }
    return books, nil
}

// WriteBooks writes the book data to the JSON file
func WriteBooks(filename string, books []models.Book) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    bytes, err := json.Marshal(books)
    if err != nil {
        return err
    }

    _, err = file.Write(bytes)
    return err
}
