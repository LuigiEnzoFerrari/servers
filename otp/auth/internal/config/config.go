package config

import (
	"fmt"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	RabbitMQ RabbitMQConfig
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

type RabbitMQConfig struct {
    User     string
    Password string
    Host     string
    Port     string
    VHost    string
}

type ServerConfig struct {
	Port string
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", ""),
			Port:     getEnv("DB_PORT", ""),
			SSLMode:  getEnv("DB_SSL_MODE", ""),
		},
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
	}
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
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
