package dto

import "time"

type DashboardSummaryResponse struct {
    UserID           string    `json:"user_id"`
    AvailableBalance float64   `json:"available_balance"`
    Currency         string    `json:"currency"`
    Status           string    `json:"status"`
    LastUpdated      time.Time `json:"last_updated"` 
    BlockedAmount    float64   `json:"blocked_amount"`
}