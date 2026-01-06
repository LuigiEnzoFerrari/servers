package service

import (
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/dto"
)

type DashboardService struct {
	
}

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}

func (s *DashboardService) GetDashboardSummary() (*dto.DashboardSummaryResponse, error) {

	response := dto.DashboardSummaryResponse{
		UserID:           "12345",
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