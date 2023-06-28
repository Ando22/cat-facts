# Use a Go base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code to the container
COPY . .

# Build the Go application
RUN go build -o app

# Expose the port on which the application will listen
EXPOSE 8000

# Set the command to run the application
CMD ["./app"]
