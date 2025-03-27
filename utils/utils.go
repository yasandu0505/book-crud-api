package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadBooks(filename string) ([]Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var books []Book
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &books)
	return books, nil
}
