package domain

import "time"

type DashboardSummary struct {
	Orders []Order
	Wallet Wallet
	User User
}

type Order struct {
	OrderID     string     
	Status      string      
	CreatedAt   time.Time   
	TotalAmount float64     
	Currency    string      
	Items       []OrderItem 
}

type OrderItem struct {
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

type Wallet struct {
	AvailableBalance float64 
	Currency string 
	Status WalletStatus 
	LastUpdated time.Time 
	BlockedAmount float64 
}

type User struct {
	UserID string
	Status string 
	CreatedAt time.Time
	LastName string
	Email string
	Phone string
	AvatarURL string
}