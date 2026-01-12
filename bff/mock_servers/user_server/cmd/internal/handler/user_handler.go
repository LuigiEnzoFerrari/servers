package handler

import (
	"net/http"
	"time"

	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/user_server/cmd/internal/handler/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUsersByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	user := mockUser()
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, user)
}

func mockUser() *dto.GetUsersByUserIDResponse {
	return &dto.GetUsersByUserIDResponse{
		UserID:    "12345",
		Status:    "ACTIVE",
		CreatedAt: time.Now(),
		PersonalInfo: dto.PersonalInfo{
			FirstName:       "John",
			LastName:        "Doe",
			Email:           "john.doe@example.com",
			Phone:           "+15550199876",
			AvatarURL:       "https://cdn.store.com/avatars/user_12345.png",
			IsEmailVerified: true,
		},
		Preferences: dto.Preferences{
			Language:       "en-US",
			Currency:       "USD",
			MarketingOptIn: false,
			Notifications: dto.NotificationPreferences{
				Email: true,
				SMS:   false,
				Push:  true,
			},
		},
		Addresses: []dto.Address{
			{
				AddressID:         "addr_01",
				Label:             "Home",
				RecipientName:     "John Doe",
				Street:            "123 Maple Avenue, Apt 4B",
				City:              "Springfield",
				State:             "IL",
				ZipCode:           "62704",
				Country:           "US",
				IsDefaultShipping: true,
				IsDefaultBilling:  true,
				Type:              "RESIDENTIAL",
			},
			{
				AddressID:         "addr_02",
				Label:             "Office",
				RecipientName:     "John Doe (c/o Tech Corp)",
				Street:            "456 Innovation Blvd",
				City:              "Chicago",
				State:             "IL",
				ZipCode:           "60601",
				Country:           "US",
				IsDefaultShipping: false,
				IsDefaultBilling:  false,
				Type:              "COMMERCIAL",
			},
		},
	}
}
