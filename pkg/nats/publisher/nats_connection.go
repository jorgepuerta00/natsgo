package publisher

import "github.com/nats-io/nats.go"

type INATSConn interface {
	Publish(subject string, data []byte) error
}

type NATSConn struct {
	conn *nats.Conn
}

func NewNATSConn(conn *nats.Conn) INATSConn {
	return &NATSConn{conn: conn}
}

func (c *NATSConn) Publish(subject string, data []byte) error {
	return c.conn.Publish(subject, data)
}
