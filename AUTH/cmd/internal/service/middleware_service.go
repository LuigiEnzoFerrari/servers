package service

import (
	"errors"
)

type middlewareService struct {
	jwtRepo JwtRepository
}

func NewMiddlewareService(jwtRepo JwtRepository) *middlewareService {
	return &middlewareService{jwtRepo: jwtRepo}
}

func (s *middlewareService) ValidateToken(tokenString string) (string, error) {
	claims, err := s.jwtRepo.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	return username, nil
}