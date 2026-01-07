package infrastructure

import (
	"time"
)

type WalletStatus int

const (
	WalletStatusUnspecified WalletStatus = iota
	WalletStatusActive
	WalletStatusSuspended
	WalletStatusClosed
)

type GetUserBalanceResponse struct {
	UserID string `json:"user_id"`
	AvailableBalance float64 `json:"available_balance"`
	Currency string `json:"currency"`
	Status WalletStatus `json:"status"`
	LastUpdated time.Time `json:"last_updated"`
	BlockedAmount float64 `json:"blocked_amount"`
}