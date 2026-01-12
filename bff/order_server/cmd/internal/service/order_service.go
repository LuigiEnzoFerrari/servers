package service

import (
	"context"
	"github.com/LuigiEnzoFerrari/servers/bff/order_server/cmd/internal/dto"
	"time"
)


type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) GetOrdersByUserID(ctx context.Context, userID string) (*dto.GetOrdersByUserIDResponse, error) {
	return &dto.GetOrdersByUserIDResponse{
		Count: 3,
		Data: []dto.Order{
			{
				OrderID:     "ord_998",
				Status:      "PENDING",
				CreatedAt:   time.Now(),
				TotalAmount: 120.50,
				Currency:    "USD",
				Items: []dto.OrderItem{
					{
						ProductID: "prod_555",
						Quantity:  2,
						UnitPrice: 60.25,
					},
				},
			},
			{
				OrderID:     "ord_887",
				Status:      "COMPLETED",
				CreatedAt:   time.Now(),
				TotalAmount: 45.00,
				Currency:    "USD",
				Items: []dto.OrderItem{
					{
						ProductID: "prod_101",
						Quantity:  1,
						UnitPrice: 45.00,
					},
				},
			},
			{
				OrderID:     "ord_665",
				Status:      "CANCELLED",
				CreatedAt:   time.Now(),
				TotalAmount: 200.00,
				Currency:    "USD",
				Items: []dto.OrderItem{
					{
						ProductID: "prod_101",
						Quantity:  1,
						UnitPrice: 100.00,
					},
					{
						ProductID: "prod_102",
						Quantity:  1,
						UnitPrice: 100.00,
					},
				},
			},
		},
	}, nil
}
