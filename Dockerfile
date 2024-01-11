# Use an image with Go 1.20 or later for the build stage
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.* ./
RUN go mod download

# Copy the application source code to the container
# You need to copy the entire project since your main.go might depend on other packages within your project
COPY . .

# Build the Go application - adjust the path to where main.go is located within your project structure
RUN CGO_ENABLED=0 go build -o ordermanager ./cmd

# Create a minimal runtime image
FROM alpine:latest

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the built binary from the build stage to the runtime image
COPY --from=build /app/ordermanager .

# Expose the port your application listens on (adjust if necessary)
EXPOSE 8080

# Command to run your application
CMD ["./ordermanager"]
