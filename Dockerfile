# Use the official Go image
FROM golang:1.22-alpine

# Install make and postgresql-client
RUN apk add --no-cache make postgresql-client gcc musl-dev curl

# Install the Go migration tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Ensure the Go bin directory is in PATH
ENV PATH=$PATH:/go/bin

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to take advantage of caching
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the entire source code into the container
COPY . .

# Specify the port that the application will use
EXPOSE 8081

# Run the application when the container starts
CMD ["go", "run", "/app/main.go"]
