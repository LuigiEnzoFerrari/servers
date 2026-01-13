package service

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"strings"

	"github.com/LuigiEnzoFerrari/servers/auth/cmd/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

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

type AuthRepository interface {
	Save(auth *domain.Auth) error
	FindByUsername(username string) (*domain.Auth, error)
}

type JwtRepository interface {
	ValidateToken(tokenString string) (jwt.MapClaims, error)
	GenerateToken(username string) (string, error)
}

type AuthService struct {
	repo AuthRepository
	jwtRepo JwtRepository
}

func NewAuthService(repo AuthRepository, jwtRepo JwtRepository) *AuthService {
	return &AuthService{repo: repo, jwtRepo: jwtRepo}
}

func (s *AuthService) SignUp(ctx context.Context, password string, username string) (*domain.Auth, error) {
	encodedHash, err := createHash(password, defaultParams)
	if err != nil {
		return nil, err
	}

	auth := &domain.Auth{
		Username:     username,
		PasswordHash: encodedHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repo.Save(auth); err != nil {
		return nil, err
	}
	return auth, nil
}

func (s *AuthService) Login(ctx context.Context, password string, username string) (*domain.JwtToken, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if ok, err := ComparePasswordAndHash(password, user.PasswordHash); err != nil || !ok {
		return nil, errors.New("invalid credentials")
	}

	tokenString, err := s.jwtRepo.GenerateToken(user.Username)

	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &domain.JwtToken{
		Token: tokenString,
	}, nil

}

func (s *AuthService) GenerateToken(ctx context.Context, username string) (string, error) {
	return s.jwtRepo.GenerateToken(username)
}


func (s *AuthService) Protected(ctx context.Context) {

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

func ComparePasswordAndHash(password, encodedHash string) (bool, error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (p *argon2Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("invalid hash format")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version")
	}

	p = &argon2Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}