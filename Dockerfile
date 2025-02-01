# Automatically detect the architecture
FROM --platform=$TARGETPLATFORM golang:1.23 AS builder

# Enable CGO and set appropriate OS/Architecture
ENV CGO_ENABLED=1 GOOS=linux CC=gcc

# Set working directory
WORKDIR /app

# Install required dependencies
RUN apt-get update && apt-get install -y gcc musl-dev sqlite3 ca-certificates

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Copy GitHub's root CA certificate to container's trusted certificate store
COPY github-root-ca.crt /usr/local/share/ca-certificates/github-root-ca.crt

# Update CA certificates inside the container to include GitHub's root CA
RUN update-ca-certificates

# Build the Go binary
RUN go build -o app

# Use a minimal runtime image
FROM debian:stable-slim

# Set working directory
WORKDIR /root/

# Install SQLite runtime dependencies and CA certificates
RUN apt-get update && apt-get install -y sqlite3 ca-certificates

# Copy compiled Go binary
COPY --from=builder /app/app .

# Copy entrypoint script and make it executable
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Expose application port
EXPOSE 8080

# Run the entrypoint script
CMD ["/entrypoint.sh"]