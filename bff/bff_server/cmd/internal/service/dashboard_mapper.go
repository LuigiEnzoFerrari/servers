package service

import (
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/domain"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/dto"
	"github.com/LuigiEnzoFerrari/servers/bff/bff_server/cmd/internal/infrastructure"
)

func mapOrdersResponseToOrders(response *infrastructure.GetOrdersByUserIDResponse) []domain.Order {
	orders := make([]domain.Order, len(response.Data))
	for i, order := range response.Data {
		orders[i] = domain.Order{
			OrderID:           order.OrderID,
			TotalAmount:		 order.TotalAmount,
			Currency:         order.Currency,
			Status:           order.Status,
			CreatedAt:        order.CreatedAt,
			Items:            mapOrderItems(order.Items),
		}
	}
	return orders
}

func mapOrderItems(items []infrastructure.OrderItem) []domain.OrderItem {
	orderItems := make([]domain.OrderItem, len(items))
	for i, item := range items {
		orderItems[i] = domain.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
	}
	return orderItems
}


func mapWalletResponseToWallet(response *infrastructure.GetUserBalanceResponse) domain.Wallet {
	return domain.Wallet{
		AvailableBalance: response.AvailableBalance,
		Currency:         response.Currency,
		Status:           domain.WalletStatus(response.Status),
		LastUpdated:      response.LastUpdated,
		BlockedAmount:    response.BlockedAmount,
	}
}

func mapUserResponseToUser(response *infrastructure.GetUserByUserIDResponse) domain.User {
	return domain.User{
		UserID:           response.UserID,
		Status:           response.Status,
		CreatedAt:        response.CreatedAt,
		LastName:         response.PersonalInfo.LastName,
		Email:            response.PersonalInfo.Email,
		Phone:            response.PersonalInfo.Phone,
		AvatarURL:        response.PersonalInfo.AvatarURL,
	}
}

func mapDashboardSummaryToDashboardSummaryResponse(summary *domain.DashboardSummary) *dto.DashboardSummaryResponse {
	return &dto.DashboardSummaryResponse{
		Orders: mapOrdersToOrderResponse(summary.Orders),
		Wallet: mapWalletToWalletResponse(summary.Wallet),
		User:   mapUserToUserResponse(summary.User),
	}
}

func mapOrdersToOrderResponse(order []domain.Order) []dto.Order {
	orders := make([]dto.Order, len(order))
	for i, order := range order {
		orders[i] = dto.Order{
			OrderID:     order.OrderID,
			Status:      order.Status,
			CreatedAt:   order.CreatedAt,
			TotalAmount: order.TotalAmount,
			Currency:    order.Currency,
			Items:       mapOrderItemsToOrderItemsResponse(order.Items),
		}
	}
	return orders
}

func mapOrderItemsToOrderItemsResponse(items []domain.OrderItem) []dto.OrderItem {
	orderItems := make([]dto.OrderItem, len(items))
	for i, item := range items {
		orderItems[i] = dto.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}
	}
	return orderItems
}

func mapWalletToWalletResponse(wallet domain.Wallet) dto.Wallet {
	return dto.Wallet{
		AvailableBalance: wallet.AvailableBalance,
		Currency:         wallet.Currency,
		Status:           mapWalletStatusToWalletStatusResponse(wallet.Status),
		LastUpdated:      wallet.LastUpdated,
		BlockedAmount:    wallet.BlockedAmount,
	}
}

func mapWalletStatusToWalletStatusResponse(walletStatus domain.WalletStatus) dto.WalletStatus {
	switch walletStatus {
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

func mapUserToUserResponse(user domain.User) dto.User {
	return dto.User{
		UserID:           user.UserID,
		Status:           user.Status,
		CreatedAt:        user.CreatedAt,
		LastName:         user.LastName,
		Email:            user.Email,
		Phone:            user.Phone,
		AvatarURL:        user.AvatarURL,
	}
}
