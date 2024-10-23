# Use the official Golang image as a base image
FROM golang:1.23.2-alpine3.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o book-management

# Expose port 8080 for the Gin server
EXPOSE 8080

# Command to run the application
CMD ["./book-management"]
