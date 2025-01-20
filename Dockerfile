# Use the specified base image
FROM mcr.microsoft.com/devcontainers/base:alpine-3.20

# Set working directory inside the container
WORKDIR /app

# Copy the setup script into the container and execute it
COPY .devcontainer/setup.sh /setup.sh
RUN chmod +x /setup.sh && /setup.sh

# Copy Go modules files separately for better caching
COPY go.mod ./
RUN go mod download

# Copy the entire project into the container
COPY . /app

# Build the Go application
RUN go build -o main .

# Expose the necessary port
EXPOSE 8003

# Command to run the compiled binary
CMD ["./main"]
