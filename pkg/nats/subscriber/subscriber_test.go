package subscriber

import (
	"errors"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNATSConn struct {
	mock.Mock
}

func (m *MockNATSConn) SubscribeSync(subject string) (*nats.Subscription, error) {
	args := m.Called(subject)
	result, _ := args.Get(0).(*nats.Subscription)
	return result, args.Error(1)
}

func TestSubscriber_SubscribeToQueue(t *testing.T) {
	// Arrange
	mockConn := new(MockNATSConn)
	subscriber := NewSubscriber(mockConn)
	subject := "testSubject"
	mockSubscription := new(nats.Subscription)

	// Act
	mockConn.On("SubscribeSync", subject).Return(mockSubscription, nil).Once()
	err := subscriber.SubscribeToSubject(subject)

	// Assert
	assert.NoError(t, err, "SubscribeToSubject should not return an error")
	mockConn.AssertExpectations(t)
}

func TestSubscriber_SubscribeToQueue_Error(t *testing.T) {
	// Arrange
	mockConn := new(MockNATSConn)
	subscriber := NewSubscriber(mockConn)
	subject := "testSubject"
	expectedErr := errors.New("subscribe error")

	// Act
	mockConn.On("SubscribeSync", subject).Return(nil, expectedErr).Once()
	err := subscriber.SubscribeToSubject(subject)

	// Assert
	assert.Error(t, err, "SubscribeToSubject should return an error")
	assert.Equal(t, expectedErr, err, "Returned error should match expected error")
	mockConn.AssertExpectations(t)
}
