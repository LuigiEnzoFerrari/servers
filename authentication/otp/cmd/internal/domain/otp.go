package domain

import (
	"encoding/json"
	"time"
)

type PasswordForgotEvent struct {
	Username string `json:"username"`
}

type SmtpService interface {
	SendOTPEmail(toEmail string, otpCode string) error
}

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   json.RawMessage       `json:"payload"`
	OccurredAt time.Time `json:"occurred_at"`
}
