package config

import "os"

type Config struct {
    WalletClient GrpcClient
    OrderClient HttpClient
    UserClient HttpClient
}

type HttpClient struct {
    port string
    host string
    version string
    resource string
}

type GrpcClient struct {
    port string
    host string
}

func NewConfig() *Config {
    return &Config{
        WalletClient: GrpcClient{
            port: os.Getenv("WALLET_PORT"),
            host: os.Getenv("WALLET_HOST"),
        },
        OrderClient: HttpClient{
            port: os.Getenv("ORDER_PORT"),
            host: os.Getenv("ORDER_HOST"),
            version: os.Getenv("ORDER_VERSION"),
            resource: os.Getenv("ORDER_RESOURCE"),
        },
        UserClient: HttpClient{
            port: os.Getenv("USER_PORT"),
            host: os.Getenv("USER_HOST"),
            version: os.Getenv("USER_VERSION"),
            resource: os.Getenv("USER_RESOURCE"),
        },
    }
}

func (c *GrpcClient) GetGrpcAddress() string {
    return c.host + ":" + c.port
}

func (c *HttpClient) GetHttpAddress() string {
    return "http://" + c.host + ":" + c.port + "/" + c.version + "/" + c.resource
}

    