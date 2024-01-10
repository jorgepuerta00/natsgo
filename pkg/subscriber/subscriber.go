package subscriber

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

	// Run subscriber
	runSubscriber(nc)

	// Keep the program running
	runtime.Goexit()
}

func runSubscriber(nc *nats.Conn) {
	subject := "test.subject"

	// Subscribe to a subject
	_, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf("Received message on subject '%s': %s", msg.Subject, string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}
