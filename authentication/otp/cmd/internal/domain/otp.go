package domain

import (
	"context"
	"encoding/json"
	"time"
)


type OptService interface {
	SendOTPEmail(ctx context.Context, body []byte) error
	VerifyOTP(ctx context.Context, email string, otpCode string) error

}

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

type OptRepository interface {
	Save(ctx context.Context, key string, otpCode string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}