package service

import (
	"context"

	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/order_server/cmd/internal/dto"
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) GetOrdersByUserID(ctx context.Context, userID string) (*dto.GetOrdersByUserIDResponse, error) {

}
