package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareService interface {
	ValidateToken(tokenString string) (string, error)
}

type AuthMiddleware struct {
	middlewareService MiddlewareService
}

func NewAuthMiddleware(middlewareService MiddlewareService) *AuthMiddleware {
	return &AuthMiddleware{middlewareService: middlewareService}
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}
		username, err := m.middlewareService.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.Set("username", username)
		c.Next()
	}
}