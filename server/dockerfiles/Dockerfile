# Use the official Golang image as a parent image
FROM golang:1.21.0

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and their dependencies list
COPY go.mod go.sum ./

# Download and install any dependencies
RUN go mod download

# Copy the rest of your application code to the container
COPY . .

# Build the WebSocket server binary
RUN go build -o app ./cmd

# Expose the port that your WebSocket server listens on
EXPOSE 8080

# Command to start your WebSocket server
CMD ["./app"]
