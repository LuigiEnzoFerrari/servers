package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"github.com/LuigiEnzoFerrari/servers/otp/auth/cmd/internal/domain"
)

type AuthService struct {
	authRepository domain.AuthRepository
}

func NewAuthService(authRepository domain.AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

type argon2Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var defaultParams = &argon2Params{
	memory:      64 * 1024,
	iterations:  1,
	parallelism: 4,
	saltLength:  16,
	keyLength:   32,
}

func (s *AuthService) Register(dto domain.AuthRequestDTO) error {
	hashPassword, err := createHash(dto.Password, defaultParams)
	if err != nil {
		return err
	}
	auth := domain.Auth{
		Email: dto.Email,
		PasswordHash: hashPassword,
	}
	return s.authRepository.Save(&auth)
}

func createHash(password string, p *argon2Params) (string, error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
