package service

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"crypto/rand"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/smtp"
)

type OptService struct {
	smtpService *smtp.MailHogService
}

func NewOptService(smtpService *smtp.MailHogService) *OptService {
	return &OptService{
		smtpService: smtpService,
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

	err = s.smtpService.SendOTP(event.Email, otp)
	if err != nil {
		log.Printf("Failed to send OTP email: %v", err)
		return err
	}
	return nil
}
