package grpc_client

import (
	pb "github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/api/proto/wallet/v1"
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/domain"
)

func MapProtoToWallet(w *pb.UserBalanceResponse) *domain.ExternalWallet {
	return &domain.ExternalWallet{
		UserID:           w.UserId,
		AvailableBalance: w.AvailableBalance,
		Currency:         w.Currency,
		Status:           domain.WalletStatus(w.Status),
		LastUpdated:      w.LastUpdated.AsTime(),
		BlockedAmount:    w.BlockedAmount,
	}
}
