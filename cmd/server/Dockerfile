# cmd/server/Dockerfile
FROM golang:1.23

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Create videos directory
RUN [ ! -d "/app/videos" ] && mkdir /app/videos || true

# Build the server binary
WORKDIR /app/cmd/server
RUN go build -o server .

# Expose ports
EXPOSE 9000
EXPOSE 50051

# Default command
CMD ["./server"]
