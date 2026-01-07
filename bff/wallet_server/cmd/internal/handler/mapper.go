package handler

import (
	pb "github.com/LuigiEnzoFerrari/servers/bff/wallet_server/api/proto/v1"
	"github.com/LuigiEnzoFerrari/servers/bff/wallet_server/cmd/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapWalletBalanceToProto(u *domain.Wallet) *pb.UserBalanceResponse {
    return &pb.UserBalanceResponse{
        UserId:    u.UserID,
        AvailableBalance: u.AvailableBalance,
        Currency: u.Currency,
        Status: mapWalletStatusToProto(u.Status),
        LastUpdated: timestamppb.New(u.LastUpdated),
        BlockedAmount: u.BlockedAmount,
    }
}

func mapWalletStatusToProto(r domain.WalletStatus) pb.UserStatus {
    switch r {
    case domain.WalletStatusActive:
        return pb.UserStatus_USER_STATUS_ACTIVE
    case domain.WalletStatusSuspended:
        return pb.UserStatus_USER_STATUS_SUSPENDED
    case domain.WalletStatusClosed:
        return pb.UserStatus_USER_STATUS_CLOSED
    default:
        return pb.UserStatus_USER_STATUS_UNSPECIFIED
    }
}