package infrastructure

import (
	pb "github.com/LuigiEnzoFerrari/servers/bff/bff_server/api/proto/wallet/v1"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/service"
)

func MapProtoToWallet(w *pb.UserBalanceResponse) *service.GetUserBalanceResponse {
	return &service.GetUserBalanceResponse{
		UserID:           w.UserId,
		AvailableBalance: w.AvailableBalance,
		Currency:         w.Currency,
		Status:           service.WalletStatus(w.Status),
		LastUpdated:      w.LastUpdated.AsTime(),
		BlockedAmount:    w.BlockedAmount,
	}
}