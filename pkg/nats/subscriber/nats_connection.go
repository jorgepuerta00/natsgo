package subscriber

import "github.com/nats-io/nats.go"

type INATSConn interface {
	SubscribeSync(subject string) (*nats.Subscription, error)
}

type NATSConn struct {
	conn *nats.Conn
}

func NewNATSConn(conn *nats.Conn) INATSConn {
	return &NATSConn{conn: conn}
}

func (c *NATSConn) SubscribeSync(subject string) (*nats.Subscription, error) {
	return c.conn.SubscribeSync(subject)
}
