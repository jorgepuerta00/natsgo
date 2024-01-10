
# Demo: NATS with Golang

This demo illustrates the basic usage of NATS with Go (Golang) for event-driven communication. It includes a publisher that sends messages and a subscriber that listens for those messages.

## Prerequisites

- Go (Golang) installed on your system.
- Access to a running NATS server. You can [download and run a NATS server](https://nats.io/download/) locally.

## Installation

First, install the NATS Go client:

```bash
go get github.com/nats-io/nats.go/
```

## Components

There are two main components in this demo:

1. **Publisher**: A simple Go program that publishes messages to a NATS subject.
2. **Subscriber**: A Go program that subscribes to a NATS subject and listens for messages.

## Step-by-Step Guide

### Step 1: Create the Publisher

Create a file named `publisher.go` with the following content:

```go
package main

import (
    "log"
    "github.com/nats-io/nats.go"
)

func main() {
    // Connect to a NATS server
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

    // Publish a message
    subject := "example.subject"
    message := "Hello, NATS!"
    err = nc.Publish(subject, []byte(message))
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Published message on subject '%s': %s", subject, message)
}
```

### Step 2: Create the Subscriber

Create another file named `subscriber.go` with the following content:

```go
package main

import (
    "log"
    "runtime"
    "github.com/nats-io/nats.go"
)

func main() {
    // Connect to a NATS server
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatal(err)
    }
    defer nc.Close()

    // Subscribe to a subject
    subject := "example.subject"
    _, err = nc.Subscribe(subject, func(msg *nats.Msg) {
        log.Printf("Received message on subject '%s': %s", msg.Subject, string(msg.Data))
    })
    if err != nil {
        log.Fatal(err)
    }

    // Keep the program running
    runtime.Goexit()
}
```

### Step 3: Run the Demo

1. Start the subscriber:

   ```bash
   go run subscriber.go
   ```

2. In a separate terminal, start the publisher:

   ```bash
   go run publisher.go
   ```

The subscriber will receive and display the message sent by the publisher.

## Conclusion

This basic demo showcases how to set up a publisher and a subscriber using NATS with Go. It demonstrates the fundamental concepts of event-driven communication in a distributed system. For more advanced usage, refer to the [NATS documentation](https://docs.nats.io/).

