# Use an official Go image as the base image
FROM golang:1.24.1

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod .
RUN go mod tidy

# Copy the rest of the application code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port that your app runs on
EXPOSE 8080

# Command to run the app
CMD ["./main"]
