package service

import (
	"context"

	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/domain"

	"golang.org/x/sync/errgroup"
)

type OrderGateway interface {
	GetOrdersByUserID(ctx context.Context, userID string) ([]domain.ExternalOrder, error)
}

type UserGateway interface {
	GetUserByUserID(ctx context.Context, userID string) (*domain.ExternalUser, error)
}

type WalletGateway interface {
	GetBalance(ctx context.Context, userID string) (*domain.ExternalWallet, error)
}

type DashboardService struct {
	orderGateway  OrderGateway
	userGateway   UserGateway
	walletGateway WalletGateway
}

func NewDashboardService(
	orderGateway OrderGateway,
	userGateway UserGateway,
	walletGateway WalletGateway,
) *DashboardService {
	return &DashboardService{
		orderGateway:  orderGateway,
		userGateway:   userGateway,
		walletGateway: walletGateway,
	}
}

func (s *DashboardService) GetDashboardSummary(ctx context.Context, userID string) (*domain.DashboardSummary, error) {
	g, ctx := errgroup.WithContext(ctx)

	var (
		ordersResp  []domain.ExternalOrder
		userResp    *domain.ExternalUser
		balanceResp *domain.ExternalWallet
	)

	g.Go(func() error {
		var err error
		ordersResp, err = s.orderGateway.GetOrdersByUserID(ctx, userID)
		return err
	})

	g.Go(func() error {
		var err error
		userResp, err = s.userGateway.GetUserByUserID(ctx, userID)
		return err
	})

	g.Go(func() error {
		var err error
		balanceResp, err = s.walletGateway.GetBalance(ctx, userID)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	dashboardSummary := domain.DashboardSummary{
		Orders: ordersResp,
		User:   *userResp,
		Wallet: *balanceResp,
	}

	return &dashboardSummary, nil
}
