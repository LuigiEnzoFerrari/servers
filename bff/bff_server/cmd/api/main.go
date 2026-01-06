package main

import (
	"github.com/gin-gonic/gin"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/handler"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/infrastructure"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/service"
)



func main() {
	httpOrderGateway := infrastructure.NewHttpOrderGateway("http://localhost:8081/api/v1")
	dashboardService := service.NewDashboardService(httpOrderGateway)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	server := gin.Default()
	group := server.Group("/api/v1")
	group.GET("/dashboard-summary", dashboardHandler.GetDashboardSummary)
	
	server.Run(":8080")
}
