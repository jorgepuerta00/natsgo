package service

import (
	"testing"

	"github.com/jorgepuerta00/natsgo/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) PublishEvent(event model.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

type OrderServiceTestSuite struct {
	suite.Suite
	service *OrderService
	mockPub *MockPublisher
}

func (suite *OrderServiceTestSuite) SetupTest() {
	logger := logrus.New()
	suite.mockPub = &MockPublisher{}
	suite.service = NewOrderService(suite.mockPub, logger)
}

func (suite *OrderServiceTestSuite) TestCreateOrder() {
	// Arrange
	event := model.Event{
		OrderID: "123",
		Product: "TestProduct",
		Price:   10.5,
	}

	suite.mockPub.On("PublishEvent", event).Return(nil).Once()

	// Act
	err := suite.service.CreateOrder(event)

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockPub.AssertExpectations(suite.T())
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}
