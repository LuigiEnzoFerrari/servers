package service

import (
	"context"

	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/dto"
)

type OrderGateway interface {
	GetOrdersByUserID(ctx context.Context, userID string) (*GetOrdersByUserIDResponse, error)
}

type UserGateway interface {
	GetUsersByUserID(ctx context.Context, userID string) (*GetUserByUserIDResponse, error)
}

type WalletGateway interface {
	GetBalance(ctx context.Context, userID string) (*GetUserBalanceResponse, error)
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

	ordersResponse, err := s.orderGateway.GetOrdersByUserID(context.Background(), "12345")
	if err != nil {
		return nil, err
	}

	orders := mapOrdersResponseToOrders(ordersResponse)

	usersResponse, err := s.userGateway.GetUsersByUserID(context.Background(), "12345")
	if err != nil {
		return nil, err
	}
	users := mapUserResponseToUser(usersResponse)
	
	balance, err := s.walletGateway.GetBalance(context.Background(), "12345")
	if err != nil {
		return nil, err
	}
	wallet := mapWalletResponseToWallet(balance)

	dashboardSummary := domain.DashboardSummary{
		Orders: orders,
		User: users,
		Wallet: wallet,
	}

	response := mapDashboardSummaryToDashboardSummaryResponse(&dashboardSummary)

	return response, nil
}

func (s *DashboardService) UpdateSomething(request *dto.UpdateSomethingRequest) (*dto.UpdateSomethingResponse, error) {

	something := dto.UpdateSomethingResponse{
		Something: request.Something,
	}

	return &something, nil
	
}