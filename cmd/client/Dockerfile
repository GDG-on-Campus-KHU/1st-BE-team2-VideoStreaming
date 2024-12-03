# cmd/client/Dockerfile
FROM golang:1.23

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the client binary
WORKDIR /app/cmd/client
RUN go build -o client .

# Default command
CMD ["./client"]
