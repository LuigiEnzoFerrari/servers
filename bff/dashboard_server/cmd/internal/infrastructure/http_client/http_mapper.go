package http_client

import "github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/domain"

func MapGetOrdersByUserIDResponseToExternalOrder(response GetOrdersByUserIDResponse) []domain.ExternalOrder {
	var externalOrders []domain.ExternalOrder
	for _, order := range response.Data {
		externalOrders = append(externalOrders, *MapOrderToExternalOrder(order))
	}
	return externalOrders
}

func MapOrderToExternalOrder(order Order) *domain.ExternalOrder {
	return &domain.ExternalOrder{
		UserID:      order.OrderID,
		Currency:    order.Currency,
		Status:      MapOrderStatusToExternalOrderStatus(order.Status),
		CreatedAt:   order.CreatedAt,
		TotalAmount: order.TotalAmount,
		Items:       MapOrderItemsToExternalOrderItems(order.Items),
	}
}

func MapOrderStatusToExternalOrderStatus(status string) domain.OrderStatus {
	switch status {
	case "PENDING":
		return domain.OrderStatusPending
	case "COMPLETED":
		return domain.OrderStatusCompleted
	case "CANCELLED":
		return domain.OrderStatusCancelled
	default:
		panic("invalid order status")
	}
}

func MapOrderItemsToExternalOrderItems(items []OrderItem) []domain.ExternalOrderItem {
	var externalOrderItems []domain.ExternalOrderItem
	for _, item := range items {
		externalOrderItems = append(externalOrderItems, MapOrderItemToExternalOrderItem(item))
	}
	return externalOrderItems
}

func MapOrderItemToExternalOrderItem(item OrderItem) domain.ExternalOrderItem {
	return domain.ExternalOrderItem{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		UnitPrice: item.UnitPrice,
	}
}

func MapGetUserByUserIDResponseToExternalUser(response GetUserByUserIDResponse) *domain.ExternalUser {
	return &domain.ExternalUser{
		UserID:    response.UserID,
		Status:    MapUserStatusToExternalUserStatus(response.Status),
		CreatedAt: response.CreatedAt,
		LastName:  response.PersonalInfo.LastName,
		Email:     response.PersonalInfo.Email,
		Phone:     response.PersonalInfo.Phone,
		AvatarURL: response.PersonalInfo.AvatarURL,
	}
}

func MapUserStatusToExternalUserStatus(status string) domain.UserStatus {
	switch status {
	case "ACTIVE":
		return domain.UserStatusActive
	case "SUSPENDED":
		return domain.UserStatusSuspended
	case "DELETED":
		return domain.UserStatusDeleted
	default:
		panic("invalid user status")
	}
}
