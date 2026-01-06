package service

import (
	"context"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/dto"
)

type DashboardService struct {
	orderGateway domain.OrderGateway
}

func NewDashboardService(orderGateway domain.OrderGateway) *DashboardService {
	return &DashboardService{
		orderGateway: orderGateway,
	}
}

func (s *DashboardService) GetDashboardSummary() (*dto.DashboardSummaryResponse, error) {

	orders, err := s.orderGateway.GetOrdersByUserID(context.Background(), "12345")
	if err != nil {
		return nil, err
	}

	response := dto.DashboardSummaryResponse{
		UserID:           orders.Data[0].OrderID,
		AvailableBalance: 100.0,
		Currency:         "USD",
		Status:           "ACTIVE",
		LastUpdated:      time.Now(),
		BlockedAmount:    0.0,
	}
	return &response, nil
}

func (s *DashboardService) UpdateSomething(request *dto.UpdateSomethingRequest) (*dto.UpdateSomethingResponse, error) {

	something := domain.DashboardSomething{
		Something: request.Something,
	}

	return &dto.UpdateSomethingResponse{
		Something: something.Something,
	}, nil
	
}