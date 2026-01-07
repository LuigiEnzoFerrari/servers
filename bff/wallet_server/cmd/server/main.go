package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/LuigiEnzoFerrari/servers/bff/wallet_server/api/proto/v1"
	"google.golang.org/grpc"
	"github.com/LuigiEnzoFerrari/servers/bff/wallet_server/cmd/internal/service"
	"github.com/LuigiEnzoFerrari/servers/bff/wallet_server/cmd/internal/handler"
)



func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	walletService := service.NewWalletService()
	walletHandler := handler.NewWalletHandler(walletService)

	s := grpc.NewServer()
	pb.RegisterWalletServiceServer(s, walletHandler)

	fmt.Println("gRPC Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}