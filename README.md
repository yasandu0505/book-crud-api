# Book CRUD API

This project is a simple CRUD (Create, Read, Update, Delete) REST API built with GoLang. It manages book data using a JSON text file as a data persistence layer. This project demonstrates core GoLang concepts such as HTTP handling, JSON processing, and file I/O.

---

## Project Structure

```
book-crud-api/
├── README.md          # Instructions for the program  
├── main.go            # Entry point of the program
├── Dockerfile         # Docker file for containerization
├── handlers/          # Contains handlers for the API endpoints
│   └── book.go        # File to handle CRUD logic
├── models/            # Contains the Book struct
│   └── book.go
├── test/              # Contains the test for GET request
│   └── book_test.go
├── data/              # Contains your data storage (JSON file)
│   └── books.json     # JSON file to store book data
|   └── book_test.json # JSON file to store book test data
└── utils/             # Helper functions (like file handling)
    └── utils.go

```

---

## How to Run the Project

### Prerequisites

- GoLang installed ([Download Go](https://go.dev/dl/))
- Docker installed ([Download Docker](https://www.docker.com/))

---

## Setup Instructions (Running Locally)

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

## Running the Project with Docker

### Build the Docker Image

To containerize the Book CRUD API, follow these steps:

1. Build the Docker image:

   ```bash
   docker build -t book-crud-api .
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 book-crud-api
   ```

3. The server will be running inside the container on `http://localhost:8080`.

---

## Docker Commands

Here are additional Docker commands you may find useful:

- **List running containers:**

  ```bash
  docker ps
  ```

- **Stop a running container:**

  ```bash
  docker stop <container_id>
  ```

- **Remove a container:**

  ```bash
  docker rm <container_id>
  ```

- **List Docker images:**

  ```bash
  docker images
  ```

- **Remove an image:**

  ```bash
  docker rmi <image_id>
  ```

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
  curl -X PUT http://localhost:8080/books/bb329a31-6b1e-4daa-87ee-71631aa05866 \
  -H "Content-Type: application/json" \
  -d '{
      "title": "The Great Gatsby - Updated",
      "pages": 200
  }'
  ```

- Delete a book:

  ```bash
  curl -X DELETE http://localhost:8080/books/bb329a31-6b1e-4daa-87ee-71631aa05866
  ```

---

## JSON Data Storage

The `data/books.json` file serves as the data persistence layer. This JSON file stores all book data and is read/written to during CRUD operations.

---

## License

This project is open-source and available under the [MIT License](LICENSE).

---

## Author

Developed by Yasii.🚀⚡️

