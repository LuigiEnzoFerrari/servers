package config

import (
	"fmt"
	"os"
)

type Config struct {
	Server   ServerConfig
	RabbitMQ RabbitMQConfig
	Redis    RedisConfig
	Smtp     SmtpConfig
}

type RabbitMQConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	VHost    string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

type SmtpConfig struct {
	Host     string
	Port     string
	Password string
	Sender   string
}

type ServerConfig struct {
	Port string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		RabbitMQ: RabbitMQConfig{
			User:     getEnv("RABBITMQ_USER", ""),
			Password: getEnv("RABBITMQ_PASSWORD", ""),
			Host:     getEnv("RABBITMQ_HOST", ""),
			Port:     getEnv("RABBITMQ_PORT", ""),
			VHost:    getEnv("RABBITMQ_VHOST", ""),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", ""),
			Port:     getEnv("REDIS_PORT", ""),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
		Smtp: SmtpConfig{
			Host:     getEnv("SMTP_HOST", ""),
			Port:     getEnv("SMTP_PORT", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
			Sender:   getEnv("SMTP_SENDER", ""),
		},
	}
}

func (c *RabbitMQConfig) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		c.User, c.Password, c.Host, c.Port, c.VHost)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
