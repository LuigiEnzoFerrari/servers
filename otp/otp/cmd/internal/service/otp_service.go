package service

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"log"
	"math/big"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/smtp"
)

type OptService struct {
	smtpService *smtp.MailHogService
	otpRepository domain.OptRepository
}

func NewOptService(smtpService *smtp.MailHogService, otpRepository domain.OptRepository) *OptService {
	return &OptService{
		smtpService: smtpService,
		otpRepository: otpRepository,
	}
}

func GenerateOTP(length int) (string, error) {
    const digits = "0123456789"
    
    otp := make([]byte, length)
    
    maxIndex := big.NewInt(int64(len(digits)))

    for i := 0; i < length; i++ {
        num, err := rand.Int(rand.Reader, maxIndex)
        if err != nil {
            return "", err
        }
        
        otp[i] = digits[num.Int64()]
    }

    return string(otp), nil
}

func (s *OptService) SendOTPEmail(ctx context.Context, body []byte) error {
	log.Println("Sending OTP email")
	var event domain.PasswordForgotEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("Failed to unmarshal order: %v", err)
		return err
	}

	otp, err := GenerateOTP(6)
	if err != nil {
		log.Printf("Failed to generate OTP: %v", err)
		return err
	}

	key := "otp:" + event.Email

	err = s.otpRepository.Save(ctx, key, otp)
	if err != nil {
		log.Printf("Failed to save OTP: %v", err)
		return err
	}

	err = s.smtpService.SendOTP(event.Email, otp)
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