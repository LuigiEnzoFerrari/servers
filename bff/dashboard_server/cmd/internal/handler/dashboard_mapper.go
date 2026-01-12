package handler

import (
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/handler/dto"
)

func MapDashboardSummaryToDashboardSummaryResponse(summary *domain.DashboardSummary) *dto.DashboardSummaryResponse {
	if summary == nil {
		return nil
	}

	return &dto.DashboardSummaryResponse{
		Orders: mapOrders(summary.Orders),
		Wallet: mapWallet(summary.Wallet),
		User:   mapUser(summary.User),
	}
}

func mapOrders(orders []domain.ExternalOrder) []dto.Order {
	if orders == nil {
		return nil
	}
	result := make([]dto.Order, len(orders))
	for i, order := range orders {
		result[i] = dto.Order{
			OrderID:     order.OrderID,
			Status:      mapOrderStatus(order.Status),
			CreatedAt:   order.CreatedAt,
			TotalAmount: order.TotalAmount,
			Currency:    order.Currency,
			Items:       mapOrderItems(order.Items),
		}
	}
	return result
}

func mapOrderItems(items []domain.ExternalOrderItem) []dto.OrderItem {
	if items == nil {
		return nil
	}
	result := make([]dto.OrderItem, len(items))
	for i, item := range items {
		result[i] = dto.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
	}
	return result
}

func mapOrderStatus(status domain.OrderStatus) string {
	switch status {
	case domain.OrderStatusPending:
		return "PENDING"
	case domain.OrderStatusCompleted:
		return "COMPLETED"
	case domain.OrderStatusCancelled:
		return "CANCELLED"
	default:
		return "UNKNOWN"
	}
}

func mapWallet(wallet domain.ExternalWallet) dto.Wallet {
	return dto.Wallet{
		AvailableBalance: wallet.AvailableBalance,
		Currency:         wallet.Currency,
		Status:           mapWalletStatus(wallet.Status),
		LastUpdated:      wallet.LastUpdated,
		BlockedAmount:    wallet.BlockedAmount,
	}
}

func mapWalletStatus(status domain.WalletStatus) dto.WalletStatus {
	switch status {
	case domain.WalletStatusActive:
		return dto.WalletStatusActive
	case domain.WalletStatusSuspended:
		return dto.WalletStatusSuspended
	case domain.WalletStatusClosed:
		return dto.WalletStatusClosed
	default:
		return dto.WalletStatusUnspecified
	}
}

func mapUser(user domain.ExternalUser) dto.User {
	return dto.User{
		UserID:    user.UserID,
		Status:    mapUserStatus(user.Status),
		CreatedAt: user.CreatedAt,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		AvatarURL: user.AvatarURL,
	}
}

func mapUserStatus(status domain.UserStatus) string {
	switch status {
	case domain.UserStatusActive:
		return "ACTIVE"
	case domain.UserStatusSuspended:
		return "SUSPENDED"
	case domain.UserStatusDeleted:
		return "DELETED"
	default:
		return "UNKNOWN"
	}
}
