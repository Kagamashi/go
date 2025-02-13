# Stage 1: Build the Go binary
FROM golang:1.20 AS builder

WORKDIR /app

# Copy source code
COPY . .

# Download dependencies and build the binary
RUN go mod tidy && go build -o myapp

# Stage 2: Run the app in a minimal image
FROM gcr.io/distroless/static AS runner

WORKDIR /

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .

# Run the binary
CMD ["/myapp"]
