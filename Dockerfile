# Build the Go application in a multi-stage Docker build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Create a new image from the alpine image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8081
CMD ["/app/main"]
