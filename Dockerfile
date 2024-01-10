FROM golang:alpine

# Move to working directory /go/src/app
WORKDIR /go/src/app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Command to run when starting the container
CMD CGO_ENABLED=0 go test ./... -v -cover