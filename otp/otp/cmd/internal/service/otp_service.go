package service

import "log"

type OptService struct {
	
}

func NewOptService() *OptService {
	return &OptService{}
}

func (s *OptService) GenerateOTP(email string) (string, error) {
	log.Println("Generating OTP for email: ", email)
	return "", nil
}

