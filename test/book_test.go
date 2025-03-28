package main

import (
	"book-crud-api/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleBooks_GET(t *testing.T) {
	// Create a new HTTP request (GET request)
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Assuming you pass the correct handler to serve the request
	handler := http.HandlerFunc(handlers.HandleBooks)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Trim both the response body and expected body before comparison
	expected := `[{"bookId":"1","authorId":"123","publisherId":"456","title":"Moby Dick","publicationDate":"1851-10-18","isbn":"9781503280786","pages":635,"genre":"Adventure","description":"A thrilling tale of a captain's obsessive quest to hunt the elusive white whale.","price":18.75,"quantity":4}]`
	actual := strings.TrimSpace(rr.Body.String())

	if actual != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}
}
