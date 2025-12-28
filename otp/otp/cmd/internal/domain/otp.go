package domain

import "context"

type Opt struct {
	
}

type OptService interface {
	SendOTPEmail(ctx context.Context, body []byte) error
}

type PasswordForgotEvent struct {
	Email string `json:"email"`
}

type SmtpService interface {
	SendOTPEmail(toEmail string, otpCode string) error
}
