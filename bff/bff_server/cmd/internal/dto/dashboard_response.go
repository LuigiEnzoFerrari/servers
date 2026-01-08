package dto

import "time"

type DashboardSummaryResponse struct {
	Orders []Order
	Wallet Wallet
	User User
}

type Order struct {
	OrderID     string    `json:"order_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	TotalAmount float64   `json:"total_amount"`
	Currency    string    `json:"currency"`
	Items       []OrderItem 
}

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity int `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type WalletStatus string

const (
	WalletStatusUnspecified WalletStatus = "UNSPECIFIED"
	WalletStatusActive WalletStatus = "ACTIVE"
	WalletStatusSuspended WalletStatus = "SUSPENDED"
	WalletStatusClosed WalletStatus = "CLOSED"
)

type Wallet struct {
	AvailableBalance float64 `json:"available_balance"`
	Currency string `json:"currency"`
	Status WalletStatus `json:"status"`
	LastUpdated time.Time `json:"last_updated"`
	BlockedAmount float64 `json:"blocked_amount"`
}

type User struct {
	UserID string `json:"user_id"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	AvatarURL string `json:"avatar_url"`
}
