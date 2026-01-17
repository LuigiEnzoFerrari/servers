package domain

import "errors"

var (
	InvalidOTPError = errors.New("Invalid OTP")
	OTPNotFoundError = errors.New("OTP not found")
)