package publisher

import (
	"log"

	"github.com/nats-io/nats.go"
)

// RunPublisher sends a single message to a NATS subject
func RunPublisher(nc *nats.Conn, message string) {
	subject := "test.subject"

	err := nc.Publish(subject, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published message on subject '%s': %s", subject, message)
}
