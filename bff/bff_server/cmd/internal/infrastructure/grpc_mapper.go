package infrastructure

import (
	pb "github.com/LuigiEnzoFerrari/servers/bff/bff_server/api/proto/wallet/v1"
)

func MapProtoToWallet(w *pb.UserBalanceResponse) *GetUserBalanceResponse {
	return &GetUserBalanceResponse{
		UserID:           w.UserId,
		AvailableBalance: w.AvailableBalance,
		Currency:         w.Currency,
		Status:           WalletStatus(w.Status),
		LastUpdated:      w.LastUpdated.AsTime(),
		BlockedAmount:    w.BlockedAmount,
	}
}