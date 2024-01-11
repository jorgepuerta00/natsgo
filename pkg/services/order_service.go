package service

import (
	"github.com/jorgepuerta00/natsgo/pkg/model"
	"github.com/jorgepuerta00/natsgo/pkg/nats/publisher"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	publisher publisher.Publisher
	logger    logrus.FieldLogger
}

func NewOrderService(publisher publisher.Publisher, logger logrus.FieldLogger) *OrderService {
	return &OrderService{
		publisher: publisher,
		logger:    logger,
	}
}

func (s *OrderService) CreateOrder(event model.Event) error {
	s.logger.Infof("Creating order: %+v", event)

	err := s.publisher.PublishEvent(event)
	if err != nil {
		s.logger.WithError(err).Error("Error publishing event to NATS")
		return err
	}

	s.logger.Infof("Published event to NATS: %+v", event)

	return nil
}
