package subscriber

import (
	"github.com/sirupsen/logrus"
)

type Subscriber interface {
	SubscribeToSubject(subject string) error
}

type NATSSubscriber struct {
	conn INATSConn
}

func NewSubscriber(conn INATSConn) *NATSSubscriber {
	logrus.Info("Subscriber created")
	return &NATSSubscriber{conn: conn}
}

func (s *NATSSubscriber) SubscribeToSubject(subject string) error {
	_, err := s.conn.SubscribeSync(subject)
	if err != nil {
		logrus.Errorf("Failed to subscribe to subject %s: %v", subject, err)
		return err
	}
	logrus.Infof("Subscribed to subject %s", subject)
	return nil
}
