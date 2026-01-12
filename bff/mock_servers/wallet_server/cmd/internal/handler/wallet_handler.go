package handler

import (
	"context"
	"log"
	"time"

	pb "github.com/LuigiEnzoFerrari/servers/bff/mock_servers/wallet_server/api/proto/v1"
	"github.com/LuigiEnzoFerrari/servers/bff/mock_servers/wallet_server/cmd/internal/domain"
)

type WalletUseCase interface {
	GetBalance(userID string) (*domain.Wallet, error)
}

type WalletHandler struct {
	pb.UnimplementedWalletServiceServer
	useCase WalletUseCase
}

func NewWalletHandler(useCase WalletUseCase) pb.WalletServiceServer {
	return &WalletHandler{useCase: useCase}
}

func (s *WalletHandler) GetBalance(ctx context.Context, req *pb.GetUserBalanceRequest) (*pb.UserBalanceResponse, error) {
	log.Printf("Received request for User ID: %s", req.GetUserId())

	balance, err := s.useCase.GetBalance(req.GetUserId())
	if err != nil {
		return nil, err
	}

	time.Sleep(3 * time.Second)
	return MapWalletBalanceToProto(balance), nil
}
