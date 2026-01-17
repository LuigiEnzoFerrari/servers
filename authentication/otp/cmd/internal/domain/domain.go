package domain

import (
)

type SmtpService interface {
	SendOTPEmail(toEmail string, otpCode string) error
}

