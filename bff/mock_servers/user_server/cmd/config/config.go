package config

import "os"

type Config struct {
    Server Server
}

type Server struct {
    port string
}

func NewConfig() *Config {
    return &Config{
        Server: Server{
            port: os.Getenv("PORT"),
        },
    }
}

func (s *Server) GetPort() string {
    return s.port
}

