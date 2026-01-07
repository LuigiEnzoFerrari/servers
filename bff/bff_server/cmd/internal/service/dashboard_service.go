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

type WalletGateway interface {
	GetBalance(ctx context.Context, userID string) (*infrastructure.GetUserBalanceResponse, error)
}

type DashboardService struct {
	orderGateway OrderGateway
	userGateway UserGateway
	walletGateway WalletGateway
}

func NewDashboardService(
	orderGateway OrderGateway,
	userGateway UserGateway,
	walletGateway WalletGateway,
) *DashboardService {
	return &DashboardService{
		orderGateway: orderGateway,
		userGateway: userGateway,
		walletGateway: walletGateway,
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
	
	balance, err := s.walletGateway.GetBalance(context.Background(), "12345")
	if err != nil {
		return nil, err
	}
	response := dto.DashboardSummaryResponse{
		UserID:           users.UserID,
		AvailableBalance: balance.AvailableBalance,
		Currency:         orders.Data[0].Currency,
		Status:           mapWalletStatusToDashboardStatus(balance.Status),
		LastUpdated:      time.Now(),
		BlockedAmount:    0.0,
	}
	return &response, nil
}

func mapWalletStatusToDashboardStatus(status infrastructure.WalletStatus) string {
	switch status {
	case infrastructure.WalletStatusActive:
		return "ACTIVE"
	case infrastructure.WalletStatusSuspended:
		return "SUSPENDED"
	case infrastructure.WalletStatusClosed:
		return "CLOSED"
	default:
		return "UNSPECIFIED"
	}
}

func (s *DashboardService) UpdateSomething(request *dto.UpdateSomethingRequest) (*dto.UpdateSomethingResponse, error) {

	something := dto.UpdateSomethingResponse{
		Something: request.Something,
	}

	return &something, nil
	
}