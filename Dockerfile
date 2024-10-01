# Use the official Golang image
FROM golang:1.22.5-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/server/main.go

# List the contents of the /app directory to verify the build
RUN ls -la /app

# List the contents of the /app/cmd/server directory to verify the source files
RUN ls -la /app/cmd/server

# Expose port 8080 for inter-container communication
EXPOSE 8080

# Command to run the executable
CMD ["./main"]