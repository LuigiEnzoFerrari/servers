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
			Port: os.Getenv("SERVER_PORT"),
		},
		RabbitMQ: RabbitMQConfig{
			User:     os.Getenv("RABBITMQ_USER"),
			Password: os.Getenv("RABBITMQ_PASSWORD"),
			Host:     os.Getenv("RABBITMQ_HOST"),
			Port:     os.Getenv("RABBITMQ_PORT"),
			VHost:    os.Getenv("RABBITMQ_VHOST"),
		},
		Redis: RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		Smtp: SmtpConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Sender:   os.Getenv("SMTP_SENDER"),
		},
	}
}

func (c *RabbitMQConfig) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		c.User, c.Password, c.Host, c.Port, c.VHost)
}

func (sv *ServerConfig) GetPort() string {
	return sv.Port
}