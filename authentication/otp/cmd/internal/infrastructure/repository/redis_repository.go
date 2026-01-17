package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/LuigiEnzoFerrari/servers/otp/otp/cmd/internal/domain"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Save(ctx context.Context, key string, otpCode string) error {
	if err := r.client.Set(ctx, key, otpCode, 5*time.Minute).Err(); err != nil {
		return fmt.Errorf("redis save error: %w", err)
	}
	return nil
}

func (r *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	otpCode, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", domain.OTPNotFoundError
	}
	return otpCode, nil
}

func (r *RedisRepository) Delete(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis delete error: %w", err)
	}
	return nil
}
