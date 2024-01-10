package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/jorgepuerta00/natsgo/pkg/publisher"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Send 5 random messages
	for i := 0; i < 5; i++ {
		message := "Message " + randomString(5)
		publisher.RunPublisher(nc, message)
		time.Sleep(1 * time.Second)
	}
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
