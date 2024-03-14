# Start from the official Go image to build your application
FROM golang:latest

# Env variables for CompileDaemon
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Install CompileDaemon locally
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Use CompileDaemon to monitor changes in Go files and rebuild
ENTRYPOINT CompileDaemon -build="go build -o main ." -command="./main"
