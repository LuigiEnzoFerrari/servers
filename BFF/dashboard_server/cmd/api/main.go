package main

import (
	"log"

	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/handler"
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/infrastructure/grpc_client"
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/infrastructure/http_client"
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/service"
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getGrpcClient(address string) *grpc.ClientConn {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func main() {

	config := config.NewConfig()

	grpcWalletClient := config.WalletClient
	orderClient := config.OrderClient
	userClient := config.UserClient

	conn := getGrpcClient(grpcWalletClient.GetGrpcAddress())
	defer conn.Close()
	grpcWalletGateway := grpc_client.NewGrpcWalletGateway(conn)
	httpOrderGateway := http_client.NewHttpOrderGateway(orderClient.GetHttpAddress())
	httpUserGateway := http_client.NewHttpUserGateway(userClient.GetHttpAddress())
	dashboardService := service.NewDashboardService(
		httpOrderGateway,
		httpUserGateway,
		grpcWalletGateway,
	)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	server := gin.Default()
	group := server.Group("/api/v1")
	group.GET("/dashboard-summary/:userId", dashboardHandler.GetDashboardSummary)

	server.Run(":8080")
}

// protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/wallet/v1/wallet.proto
