package service

import (
	"context"
	"log"
)

type OptService struct {
	
}

func NewOptService() *OptService {
	return &OptService{}
}

func (s *OptService) GenerateOTP(ctx context.Context, body []byte) error {
	log.Println("Generating OTP for email: ", body)
	return nil
}

