package domain

import "time"

type WalletStatus int

const (
	WalletStatusUnspecified WalletStatus = iota
	WalletStatusActive
	WalletStatusSuspended
	WalletStatusClosed
)

type Wallet struct {
	UserID string
	AvailableBalance float64
	Currency string
	Status WalletStatus
	LastUpdated time.Time
	BlockedAmount float64
}


