package subscriber

import (
	"github.com/sirupsen/logrus"
)

type Subscriber interface {
	SubscribeToSubject(subject string) error
}

type NATSSubscriber struct {
	conn   INATSConn
	logger logrus.FieldLogger
}

func NewSubscriber(conn INATSConn, logger logrus.FieldLogger) *NATSSubscriber {
	logger.Infof("Subscriber created")
	return &NATSSubscriber{
		conn:   conn,
		logger: logger,
	}
}

func (s *NATSSubscriber) SubscribeToSubject(subject string) error {
	_, err := s.conn.SubscribeSync(subject)
	if err != nil {
		s.logger.Errorf("Failed to subscribe to subject %s: %v", subject, err)
		return err
	}
	s.logger.Infof("Subscribed to subject %s", subject)
	return nil
}
