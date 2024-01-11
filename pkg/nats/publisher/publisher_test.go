package publisher

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/jorgepuerta00/natsgo/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNATSConn struct {
	mock.Mock
}

func (m *MockNATSConn) Publish(subject string, data []byte) error {
	args := m.Called(subject, data)
	return args.Error(0)
}

func init() {
	logrus.SetOutput(logrus.New().Writer())
}

func TestNATSPublisher_PublishEvent(t *testing.T) {
	// Arrange
	mockConn := new(MockNATSConn)
	publisher := NewNATSPublisher(mockConn, "testQueue", logrus.New())

	event := model.Event{
		OrderID: "123",
		Product: "TestProduct",
		Price:   10.5,
	}

	eventJSON, _ := json.Marshal(event)

	mockConn.On("Publish", "testQueue", eventJSON).Return(nil).Once()

	// Act
	err := publisher.PublishEvent(event)

	// Assert
	assert.NoError(t, err, "PublishEvent should not return an error")
	mockConn.AssertExpectations(t)
}

func TestNATSPublisher_PublishEvent_Error(t *testing.T) {
	// Arrange
	mockConn := new(MockNATSConn)
	publisher := NewNATSPublisher(mockConn, "testQueue", logrus.New())

	event := model.Event{
		OrderID: "123",
		Product: "TestProduct",
		Price:   10.5,
	}

	eventJSON, _ := json.Marshal(event)

	expectedErr := errors.New("publish error")
	mockConn.On("Publish", "testQueue", eventJSON).Return(expectedErr).Once()

	// Act
	err := publisher.PublishEvent(event)

	// Assert
	assert.Error(t, err, "PublishEvent should return an error")
	assert.Equal(t, expectedErr, err, "Returned error should match expected error")
	mockConn.AssertExpectations(t)
}
