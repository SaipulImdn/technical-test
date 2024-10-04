# Use the official Golang image as a build stage
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies (can be cached if go.mod and go.sum are not changed)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Copy the .env file into the container
COPY .env .env

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o technical-test ./main.go

# Use a minimal base image for the final stage
FROM alpine:latest

# Set environment variables for dynamic configuration
ENV APP_ENV=${APP_ENV:-production}
ENV APP_PORT=${APP_PORT:-5000}
ENV DB_HOST=${DB_HOST:-mysql-17870f33-saipulimdn.i.aivencloud.com}
ENV DB_USER=${DB_USER:-avnadmin}
ENV DB_PASSWORD=${DB_PASSWORD:-AVNS_twQ0pS7cHHYZvYPJvEW}
ENV DB_NAME=${DB_NAME:-defaultdb}
ENV DB_PORT=${DB_PORT:-15516}
ENV JWT_SECRET_KEY=${JWT_SECRET_KEY:-79f6df2c-1643-4612-8b81-3065c6471e66}

# Create a directory for the app
WORKDIR /app

# Copy the binary and .env file from the builder stage
COPY --from=builder /app/technical-test .
COPY --from=builder /app/.env .env

# Expose the port the app runs on
EXPOSE 5000

# Command to run the executable
CMD ["./technical-test"]
