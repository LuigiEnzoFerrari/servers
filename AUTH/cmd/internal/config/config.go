package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	ServerConfig ServerConfig
	DatabaseConfig DatabaseConfig
	JwtConfig      JwtConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
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
		JwtConfig: JwtConfig{
			Key: os.Getenv("JWT_KEY"),
			ExpireTime: time.Duration(jwtHour) * time.Hour,
		},
		ServerConfig: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
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

func (c *Config) ServerPort() string {
	return fmt.Sprintf("%s", c.ServerConfig.Port)
}


