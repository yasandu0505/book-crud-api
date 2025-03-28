# ⚡️Book CRUD API⚡️

This project is a simple CRUD (Create, Read, Update, Delete) REST API built with GoLang. It manages book data using a JSON text file as a data persistence layer. This project demonstrates core GoLang concepts such as HTTP handling, JSON processing, and file I/O.

---

## Project Structure

```
book-crud-api/
├── README.md          # Instructions for the program  
├── main.go            # Entry point of the program
├── Dockerfile         # Docker file for containerizing the app
├── handlers/          # Contains handlers for the API endpoints
│   └── book.go        # File to handle CRUD logic
├── models/            # Contains the Book struct
│   └── book.go
├── test/              # Contains the test for GET request
│   └── book_test.go
├── data/              # Contains your data storage (JSON files)
│   └── books.json     # JSON file to store book data for production
│   └── book_test.json # JSON file used when running tests
└── utils/             # Helper functions (like file handling)
    └── utils.go
```

---

## How to Run the Project

### Prerequisites

- GoLang installed ([Download Go](https://go.dev/dl/))
- Docker installed (if you want to run the app in a container)

---

### Setup Instructions

1. Clone the repository:

   ```bash
   git clone https://github.com/yasandu0505/book-crud-api.git
   cd book-crud-api
   ```

2. Initialize Go modules:

   ```bash
   go mod init book-crud-api
   go mod tidy
   ```

3. Run the project:

   ```bash
   go run main.go
   ```

4. The server will start on `http://localhost:8080`.

---

## Running Tests

To run tests, ensure you are using the `book_test.json` file as your test data storage.  
Follow these instructions to test the GET request handler:

1. Make sure you're in the root of the project.
2. Use Go's testing command:

   ```bash
   go test ./...
   ```

   This will run the tests located in the `test/book_test.go` file.

> **Note:** The tests will read and manipulate the `data/book_test.json` file instead of `books.json`.  
This ensures your main production data file (`books.json`) remains untouched during testing.

---

## API Endpoints

### Base URL

```
http://localhost:8080
```

### Endpoints

1. **Get All Books**

   - **GET** `/books`
   - Returns a list of all books.

2. **Get Book by ID**

   - **GET** `/books?id=<id>`
   - Returns a single book by its ID.

3. **Create a New Book**

   - **POST** `/books`
   - Request Body:
     ```json
     {
       "bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866",
       "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
       "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
       "title": "The Great Gatsby",
       "publicationDate": "1925-04-10",
       "isbn": "9780743273565",
       "pages": 180,
       "genre": "Novel",
       "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
       "price": 15.99,
       "quantity": 5
     }
     ```

4. **Update a Book**

   - **PUT** `/books?id=<id>`
   - Request Body: Same as the POST body (used to update book details).

5. **Delete a Book**

   - **DELETE** `/books?id=<id>`
   - Deletes the book with the specified ID.

---

## Running the App in Docker

1. Build the Docker image:

   ```bash
   docker build -t book-crud-api .
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 book-crud-api
   ```

3. The app will be accessible at `http://localhost:8080`.

---

## Example cURL Commands

- Get all books:

  ```bash
  curl -X GET http://localhost:8080/books
  ```

- Create a new book:

  ```bash
  curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
      "bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866",
      "title": "The Great Gatsby",
      "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
      "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
      "publicationDate": "1925-04-10",
      "isbn": "9780743273565",
      "pages": 180,
      "genre": "Novel",
      "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
      "price": 15.99,
      "quantity": 5
  }'
  ```

- Update a book:

  ```bash
  curl -X PUT http://localhost:8080/books?id=bb329a31-6b1e-4daa-87ee-71631aa05866 \
  -H "Content-Type: application/json" \
  -d '{
      "title": "The Great Gatsby - Updated",
      "pages": 200
  }'
  ```

- Delete a book:

  ```bash
  curl -X DELETE http://localhost:8080/books?id=bb329a31-6b1e-4daa-87ee-71631aa05866
  ```

---

## JSON Data Storage

- The `data/books.json` file serves as the main data persistence layer.
- The `data/book_test.json` file is used specifically during testing to avoid conflicts with production data.

---

## License

This project is open-source and available under the [MIT License](LICENSE).

---

## Author

Developed by Yasii.🚀⚡️

