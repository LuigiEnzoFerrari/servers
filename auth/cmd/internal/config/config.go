package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseConfig DatabaseConfig
	RedisConfig    RedisConfig
	JwtConfig      JwtConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

type JwtConfig struct {
	Key string
	ExpireTime time.Duration
	
}

func LoadConfig() (*Config, error) {
	jwtExpireTime := os.Getenv("JWT_EXPIRE_TIME")
	jwtHour, _ := strconv.Atoi(jwtExpireTime)
	return &Config{
		DatabaseConfig: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
		},
		RedisConfig: RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		JwtConfig: JwtConfig{
			Key: os.Getenv("JWT_KEY"),
			ExpireTime: time.Duration(jwtHour) * time.Hour,
		},
	}, nil
}

func (c *Config) DNS() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.DatabaseConfig.Host,
		c.DatabaseConfig.User,
		c.DatabaseConfig.Password,
		c.DatabaseConfig.Name,
		c.DatabaseConfig.Port,
		c.DatabaseConfig.SSLMode,
	)
}

