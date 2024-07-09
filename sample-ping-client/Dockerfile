# ========== Stage 1: Build the Go binary ==========
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy module file and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build  Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o sample-ping-client .


# ========== Stage 2: Create the final image ==========
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/sample-ping-client .

# Set timezone and install ca-certificates
RUN apk add --no-cache tzdata ca-certificates

# Expose the port
EXPOSE 8080

# Set the entrypoint
ENTRYPOINT ["/app/sample-ping-client"]
