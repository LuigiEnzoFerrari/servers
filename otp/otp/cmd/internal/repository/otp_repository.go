package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisOtpRepository struct {
	client *redis.Client
}

func NewRedisOtpRepository(client *redis.Client) *RedisOtpRepository {
	return &RedisOtpRepository{client: client}
}

func (r *RedisOtpRepository) Save(ctx context.Context, key string, otpCode string) error {
	return r.client.Set(ctx, key, otpCode, 5 * time.Minute).Err()
}

func (r *RedisOtpRepository) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisOtpRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
