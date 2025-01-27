# Dockerfile
FROM golang:1.23-alpine

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
    
# Install necessary tools
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 3001

# Run the application
CMD ["./main"]