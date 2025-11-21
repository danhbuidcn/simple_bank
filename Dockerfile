# Use the official Go image
FROM golang:1.23-alpine

# Install make and postgresql-client
RUN apk add --no-cache make postgresql-client gcc musl-dev curl

# Install golang-migrate with postgres tag
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1

# Install sqlc
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.30.0

# Install air
RUN go install github.com/air-verse/air@v1.61.7

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

# Run the application with air when the container starts
CMD ["air"]
