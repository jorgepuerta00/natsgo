package publisher

import (
	"testing"

	"github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestRunPublisher(t *testing.T) {
	// Start a NATS server
	opts := test.DefaultTestOptions
	opts.Port = -1
	srv := test.RunServer(&opts)
	defer srv.Shutdown()

	// Connect to the NATS server
	nc, err := nats.Connect(srv.ClientURL())
	assert.NoError(t, err)
	defer nc.Close()

	// Test the RunPublisher function
	assert.NotPanics(t, func() {
		RunPublisher(nc, "test message")
	}, "RunPublisher should not panic")
}
