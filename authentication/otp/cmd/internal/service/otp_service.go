package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/pkg"
)

type cacheRepository interface {
	Save(ctx context.Context, key string, otpCode string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type SmtpService interface {
	SendOTP(toEmail string, otpCode string) error
}

type OptService struct {
	smtpService SmtpService
	otpRepository cacheRepository
}

func NewOptService(smtpService SmtpService,
	otpRepository cacheRepository) *OptService {
	return &OptService{
		smtpService: smtpService,
		otpRepository: otpRepository,
	}
}

func (s *OptService) SendOTPEmail(ctx context.Context, event domain.Event) error {

	var passwordForgotEvent domain.PasswordForgotEvent

	if err := json.Unmarshal(event.Payload, &passwordForgotEvent); err != nil {
		return fmt.Errorf("failed to unmarshal password forgot event: %v", err)
	}

	otp, err := pkg.GenerateOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %v", err)
	}

	key := "otp:" + passwordForgotEvent.Username

	err = s.otpRepository.Save(ctx, key, otp)
	if err != nil {
		return fmt.Errorf("failed to save OTP: %v", err)
	}

	err = s.smtpService.SendOTP(passwordForgotEvent.Username, otp)
	if err != nil {
		return fmt.Errorf("failed to send OTP email: %v", err)
	}
	return nil
}

func (s *OptService) VerifyOTP(ctx context.Context, email string, otpCode string) error {
	key := "otp:" + email
	storedOtp, err := s.otpRepository.Get(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to get OTP: %v", err)
	}
	if storedOtp != otpCode {
		return fmt.Errorf("invalid OTP")
	}
	err = s.otpRepository.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to delete OTP: %v", err)
	}
	return nil
}