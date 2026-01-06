package domain

import (
	"context"
	"time"
)

type GetOrdersByUserIDResponse struct {
	Count int     `json:"count"`
	Data  []Order `json:"data"`
}

type Order struct {
	OrderID     string      `json:"order_id"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	TotalAmount float64     `json:"total_amount"`
	Currency    string      `json:"currency"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type OrderGateway interface {
	GetOrdersByUserID(ctx context.Context, userID string) (*GetOrdersByUserIDResponse, error)
}