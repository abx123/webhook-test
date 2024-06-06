# Use the official Golang image to build and run the application
FROM golang:1.22.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go server
RUN go build -o server .

# Command to run the server
CMD ["./server"]
