package domain

import "context"


type OptService interface {
	SendOTPEmail(ctx context.Context, body []byte) error
	VerifyOTP(ctx context.Context, email string, otpCode string) error

}

type PasswordForgotEvent struct {
	Email string `json:"email"`
}

type SmtpService interface {
	SendOTPEmail(toEmail string, otpCode string) error
}

type OptRepository interface {
	Save(ctx context.Context, key string, otpCode string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}