package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/jorgepuerta00/natsgo/pkg/model"
	"github.com/jorgepuerta00/natsgo/pkg/nats/publisher"
	"github.com/jorgepuerta00/natsgo/pkg/nats/subscriber"
	service "github.com/jorgepuerta00/natsgo/pkg/services"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var (
	natsURL     = os.Getenv("NATS_URL")
	natsSubject = os.Getenv("NATS_SUBJECT")
	httpPort    = os.Getenv("HTTP_PORT")
	logger      = logrus.New()
)

func main() {

	// Set log level
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Debug("Starting application...")

	// Initialize NATS publisher
	natsPublisher := initializeNATSPublisher()

	// Initialize NATS subscriber
	initializeNatSubscriber()

	// Initialize HTTP server and routes
	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/create-order", createOrderHandler(natsPublisher))

	// Start HTTP server
	go startHTTPServer(httpHandler)

	// Keep the application running to handle NATS messages
	logger.Debug("Application running...")

	select {}
}

func initializeNATSPublisher() publisher.Publisher {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		logger.Fatalf("Error connecting to NATS: %v", err)
	}
	defer conn.Close()

	return publisher.NewNATSPublisher(conn, natsSubject)
}

func initializeNatSubscriber() subscriber.Subscriber {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		logger.Fatalf("Error connecting to NATS: %v", err)
	}
	defer conn.Close()

	return subscriber.NewSubscriber(conn)
}

func createOrderHandler(publisher publisher.Publisher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
			return
		}

		event, err := decodeJSONRequest(r.Body)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			logger.Errorf("Error decoding JSON request: %v", err)
			return
		}

		orderService := initializeOrderService(publisher)
		if err := orderService.CreateOrder(event); err != nil {
			http.Error(w, "Error creating order", http.StatusInternalServerError)
			logger.Errorf("Error creating order: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func decodeJSONRequest(body io.Reader) (model.Event, error) {
	decoder := json.NewDecoder(body)
	var event model.Event
	if err := decoder.Decode(&event); err != nil {
		return model.Event{}, err
	}
	return event, nil
}

func initializeOrderService(publisher publisher.Publisher) *service.OrderService {
	return service.NewOrderService(publisher, logger)
}

func startHTTPServer(handler http.Handler) {
	logger.Infof("HTTP server listening on port %s", httpPort)
	http.ListenAndServe(":"+httpPort, handler)
}
