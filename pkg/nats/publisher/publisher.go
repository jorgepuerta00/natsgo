// publisher/publisher.go

package publisher

import (
	"encoding/json"

	"github.com/jorgepuerta00/natsgo/pkg/model"
	"github.com/sirupsen/logrus"
)

type Publisher interface {
	PublishEvent(event model.Event) error
}

type NATSPublisher struct {
	conn    INATSConn
	subject string
}

func NewNATSPublisher(conn INATSConn, subject string) *NATSPublisher {
	return &NATSPublisher{
		conn:    conn,
		subject: subject,
	}
}

func (p *NATSPublisher) PublishEvent(event model.Event) error {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		logrus.WithError(err).Error("Failed to marshal event")
		return err
	}

	err = p.conn.Publish(p.subject, eventJSON)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to publish message to subject %s", p.subject)
		return err
	}
	logrus.Infof("Published event to subject %s", p.subject)
	return nil
}
