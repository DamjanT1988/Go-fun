# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the application
CMD ["./main"]
