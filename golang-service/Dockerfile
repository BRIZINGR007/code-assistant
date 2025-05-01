# Use official Golang image as a builder
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/server

# Runtime image
FROM debian:bookworm-slim

# Install Git and CA certificates
RUN apt-get update && apt-get install -y \
    git \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Create a non-root user
RUN useradd -m appuser

WORKDIR /app
RUN touch .env

COPY --from=builder /app/server .

USER appuser

EXPOSE 4281

ENTRYPOINT ["./server"]
