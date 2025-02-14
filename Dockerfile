# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/todo-api

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/todo-api .

# Expose default port
EXPOSE 3000

# Run the application
CMD ["./todo-api"]
