package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/LuigiEnzoFerrari/servers/bff/mock_servers/wallet_server/api/proto/v1"
	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/wallet_server/cmd/config"
	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/wallet_server/cmd/internal/handler"
	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/wallet_server/cmd/internal/service"
	"google.golang.org/grpc"
)

func main() {
	config := config.NewConfig()
	serverPort := config.Server.GetPort()
	lis, err := net.Listen("tcp", ":" + serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	walletService := service.NewWalletService()
	walletHandler := handler.NewWalletHandler(walletService)

	s := grpc.NewServer()
	pb.RegisterWalletServiceServer(s, walletHandler)

	fmt.Println("gRPC Server is running on port " + serverPort + "...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
