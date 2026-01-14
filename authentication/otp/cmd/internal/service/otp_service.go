package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"

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
		log.Printf("Failed to unmarshal password forgot event: %v", err)
		return err
	}

	log.Println("Username: " + passwordForgotEvent.Username)

	otp, err := pkg.GenerateOTP(6)
	if err != nil {
		log.Printf("Failed to generate OTP: %v", err)
		return err
	}

	key := "otp:" + passwordForgotEvent.Username

	err = s.otpRepository.Save(ctx, key, otp)
	if err != nil {
		log.Printf("Failed to save OTP: %v", err)
		return err
	}

	log.Println("Sending OTP email")
	err = s.smtpService.SendOTP(passwordForgotEvent.Username, otp)
	if err != nil {
		log.Printf("Failed to send OTP email: %v", err)
		return err
	}
	return nil
}

func (s *OptService) VerifyOTP(ctx context.Context, email string, otpCode string) error {
	key := "otp:" + email
	storedOtp, err := s.otpRepository.Get(ctx, key)
	if err != nil {
		log.Printf("Failed to get OTP: %v", err)
		return err
	}
	if storedOtp != otpCode {
		return errors.New("invalid OTP")
	}
	err = s.otpRepository.Delete(ctx, key)
	if err != nil {
		log.Printf("Failed to delete OTP: %v", err)
		return err
	}
	return nil
}