package subscriber

import (
	"testing"
	"time"

	"github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestRunSubscriber(t *testing.T) {
	// Start a NATS server
	opts := test.DefaultTestOptions
	opts.Port = -1
	srv := test.RunServer(&opts)
	defer srv.Shutdown()

	// Connect to the NATS server
	nc, err := nats.Connect(srv.ClientURL())
	assert.NoError(t, err)
	defer nc.Close()

	// Create a channel to signal the reception of a message
	msgCh := make(chan *nats.Msg, 1)

	// Subscribe with a handler that sends the received message to the channel
	_, err = nc.Subscribe("example.subject", func(msg *nats.Msg) {
		msgCh <- msg
	})
	assert.NoError(t, err)

	// Publish a message
	testMessage := "test message"
	err = nc.Publish("example.subject", []byte(testMessage))
	assert.NoError(t, err)

	// Set a timeout for receiving the message
	select {
	case receivedMsg := <-msgCh:
		assert.Equal(t, testMessage, string(receivedMsg.Data), "The message content should match")
	case <-time.After(2 * time.Second):
		t.Fatal("Did not receive message in time")
	}
}
