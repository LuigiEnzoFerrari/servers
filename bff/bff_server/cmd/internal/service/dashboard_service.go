package service

import (
	"context"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/dto"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/infrastructure"
)

type OrderGateway interface {
	GetOrdersByUserID(ctx context.Context, userID string) (*infrastructure.GetOrdersByUserIDResponse, error)
}

type UserGateway interface {
	GetUsersByUserID(ctx context.Context, userID string) (*infrastructure.GetUsersByUserIDResponse, error)
}

type DashboardService struct {
	orderGateway OrderGateway
	userGateway UserGateway
}

func NewDashboardService(
	orderGateway OrderGateway,
	userGateway UserGateway,
) *DashboardService {
	return &DashboardService{
		orderGateway: orderGateway,
		userGateway: userGateway,
	}
}

func (s *DashboardService) GetDashboardSummary() (*dto.DashboardSummaryResponse, error) {

	orders, err := s.orderGateway.GetOrdersByUserID(context.Background(), "12345")
	if err != nil {
		return nil, err
	}

	users, err := s.userGateway.GetUsersByUserID(context.Background(), "12345")
	if err != nil {
		return nil, err
	}
	response := dto.DashboardSummaryResponse{
		UserID:           users.UserID,
		AvailableBalance: 100.0,
		Currency:         orders.Data[0].Currency,
		Status:           "ACTIVE",
		LastUpdated:      time.Now(),
		BlockedAmount:    0.0,
	}
	return &response, nil
}

func (s *DashboardService) UpdateSomething(request *dto.UpdateSomethingRequest) (*dto.UpdateSomethingResponse, error) {

	something := dto.UpdateSomethingResponse{
		Something: request.Something,
	}

	return &something, nil
	
}