package main

import (
	"log"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/handler"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/infrastructure"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/service"
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

	conn := getGrpcClient("localhost:50051")
	defer conn.Close()
	grpcWalletGateway := infrastructure.NewGrpcWalletGateway(conn)
	httpOrderGateway := infrastructure.NewHttpOrderGateway("http://localhost:8081/api/v1")
	httpUserGateway := infrastructure.NewHttpUserGateway("http://localhost:8082/api/v1")
	dashboardService := service.NewDashboardService(httpOrderGateway, httpUserGateway, grpcWalletGateway)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	server := gin.Default()
	group := server.Group("/api/v1")
	group.GET("/dashboard-summary/:userId", dashboardHandler.GetDashboardSummary)
	
	server.Run(":8080")
}
// protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/wallet/v1/wallet.proto