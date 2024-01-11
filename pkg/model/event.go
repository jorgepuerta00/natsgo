package model

type Event struct {
	OrderID string  `json:"orderId"`
	Product string  `json:"product"`
	Price   float64 `json:"price"`
}

type Result struct {
	Data Event
}

type ErrorResult struct {
	ErrorMessage string `json:"errorMessage"`
}
