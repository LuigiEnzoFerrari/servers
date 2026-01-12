package domain

import "time"

type OrderStatus int

const (
	OrderStatusPending OrderStatus = iota
	OrderStatusCompleted
	OrderStatusCancelled
)

type ExternalOrder struct {
	UserID		string
	OrderID     string     
	Status      OrderStatus      
	CreatedAt   time.Time   
	TotalAmount float64     
	Currency    string      
	Items       []ExternalOrderItem 
}

type ExternalOrderItem struct {
	ProductID string
	Quantity int
	UnitPrice float64
}

type WalletStatus int

const (
	WalletStatusUnspecified WalletStatus = iota
	WalletStatusActive
	WalletStatusSuspended
	WalletStatusClosed
)

type ExternalWallet struct {
	UserID string
	AvailableBalance float64 
	Currency string 
	Status WalletStatus 
	LastUpdated time.Time 
	BlockedAmount float64 
}

type UserStatus int

const (
	UserStatusActive UserStatus = iota
	UserStatusSuspended
	UserStatusDeleted
)

type ExternalUser struct {
	UserID string
	Status UserStatus 
	CreatedAt time.Time
	LastName string
	Email string
	Phone string
	AvatarURL string
}