# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME_PUBLISHER=publisher
BINARY_NAME_SUBSCRIBER=subscriber

# Source directories
SRC_DIR_PUBLISHER=./publisher
SRC_DIR_SUBSCRIBER=./subscriber

# Build commands
all: build_publisher build_subscriber
build_publisher:
	$(GOBUILD) -o $(BINARY_NAME_PUBLISHER) -v $(SRC_DIR_PUBLISHER)
build_subscriber:
	$(GOBUILD) -o $(BINARY_NAME_SUBSCRIBER) -v $(SRC_DIR_SUBSCRIBER)

# Run commands
run_publisher: build_publisher
	./$(BINARY_NAME_PUBLISHER)
run_subscriber: build_subscriber
	./$(BINARY_NAME_SUBSCRIBER)

# Test commands
test:
	$(GOTEST) -v ./...

# Clean up
clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME_PUBLISHER)
	rm -f $(BINARY_NAME_SUBSCRIBER)

# Get dependencies
deps:
	$(GOGET) github.com/nats-io/nats.go

.PHONY: all build_publisher build_subscriber run_publisher run_subscriber test clean deps
