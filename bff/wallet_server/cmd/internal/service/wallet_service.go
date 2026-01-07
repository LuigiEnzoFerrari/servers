package service

import (
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/wallet_server/cmd/internal/domain"
)


type WalletService struct {

}

func NewWalletService() *WalletService {
	return &WalletService{}
}

func (s *WalletService) GetBalance(userID string) (*domain.Wallet, error) {

	return &domain.Wallet{
		UserID: "12345",
		AvailableBalance: 150.75,
		Currency: "USD",
		Status: domain.WalletStatusActive,
		LastUpdated: time.Now(),
		BlockedAmount: 0.00,
	}, nil
}