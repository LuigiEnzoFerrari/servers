package service

import (
	"context"

	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/dto"

	"golang.org/x/sync/errgroup"
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


func (s *DashboardService) GetDashboardSummary(ctx context.Context, userID string) (*dto.DashboardSummaryResponse, error) {
    g, ctx := errgroup.WithContext(ctx)

    var (
        ordersResp  *GetOrdersByUserIDResponse
        userResp    *GetUserByUserIDResponse
        balanceResp *GetUserBalanceResponse
    )

    g.Go(func() error {
        var err error
        ordersResp, err = s.orderGateway.GetOrdersByUserID(ctx, userID)
        return err
    })

    g.Go(func() error {
        var err error
        userResp, err = s.userGateway.GetUsersByUserID(ctx, userID)
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
        Orders: mapOrdersResponseToOrders(ordersResp),
        User:   mapUserResponseToUser(userResp),
        Wallet: mapWalletResponseToWallet(balanceResp),
    }


    return mapDashboardSummaryToDashboardSummaryResponse(&dashboardSummary), nil
}

func (s *DashboardService) UpdateSomething(ctx context.Context, request *dto.UpdateSomethingRequest) (*dto.UpdateSomethingResponse, error) {

	something := dto.UpdateSomethingResponse{
		Something: request.Something,
	}

	return &something, nil
	
}