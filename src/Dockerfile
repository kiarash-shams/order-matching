FROM golang:1.23 AS builder

# Set the working directory inside Docker to /app
WORKDIR /app

# Copy the go.mod and go.sum files for dependency management
COPY go.* ./
RUN go mod download

# Copy the entire project from src (since Dockerfile is already in src)
COPY . .

# Build the Go application with the necessary build flags
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server ./cmd/main.go

# Final stage: using a smaller base image to run the application
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive \
    rm -rf /var/lib/apt/lists/*

# Copy the built server binary and the necessary configuration files
COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/config-docker.yml /app/config/config-docker.yml

# Optionally, copy other files like documentation
COPY --from=builder /app/docs /app/docs


# Set environment variables for running the application in Docker
ENV APP_ENV=docker
ENV PORT=${Port}

# Start the application
CMD [ "/app/server" ]
