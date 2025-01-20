# Use the specified base image
FROM mcr.microsoft.com/devcontainers/base:alpine-3.20

# Set working directory inside the container
WORKDIR /app

# Copy the setup script into the container
COPY .devcontainer/setup.sh /setup.sh

# Grant execution permissions and run the setup script
RUN chmod +x /setup.sh && /setup.sh

# Copy the entire project into the container
COPY . /app

# Build the Go application
RUN go build -o server main.go

# Expose the necessary port
EXPOSE 8003

# Command to run the compiled binary
CMD ["./main"]
