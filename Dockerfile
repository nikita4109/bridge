
# Build stage
FROM golang:1.23.3-alpine AS builder

# Add build-time labels
LABEL stage=builder

# Install necessary build tools
RUN apk add --no-cache \
    git \
    protoc \
    protobuf-dev \
    make \
    gcc \
    musl-dev

# Set working directory
WORKDIR /build

# Copy go mod files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the code
COPY . .

# Build the application with production flags
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -installsuffix cgo \
    -ldflags='-w -s -extldflags "-static"' \
    -o bridge-server ./cmd/server

# Security scan stage (optional but recommended for production)
FROM aquasec/trivy:latest AS security-scan
COPY --from=builder /build /build
RUN trivy filesystem --no-progress --exit-code 1 --severity HIGH,CRITICAL /build

# Final production stage
FROM alpine:3.18

# Add production labels
LABEL maintainer="Your Name <your.email@example.com>" \
      description="Bridge Server Production Image" \
      version="1.0.0"

# Add necessary runtime dependencies
RUN apk add --no-cache \
    ca-certificates \
    curl \
    wget \
    tzdata && \
    rm -rf /var/cache/apk/*

# Create non-root user
RUN adduser -D -u 1000 appuser

# Set working directory
WORKDIR /app

# Copy the binary and config
COPY --from=builder --chown=appuser:appuser /build/bridge-server .
COPY --from=builder --chown=appuser:appuser /build/config/config.json ./config/

# Set proper permissions
RUN chmod 550 /app/bridge-server && \
    chmod 440 /app/config/config.json

# Use non-root user
USER appuser

# Expose gRPC port
EXPOSE 50051

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --spider -q http://localhost:50051 || exit 1

# Set environment
ENV GO_ENV=production \
    TZ=UTC

# Run the binary
ENTRYPOINT ["./bridge-server"]
