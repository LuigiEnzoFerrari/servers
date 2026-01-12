package grpc_client

import (
	"context"

	pb "github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/api/proto/wallet/v1"
	"github.com/LuigiEnzoFerrari/servers/bff/dashboard_server/cmd/internal/domain"
	"google.golang.org/grpc"
)

type GrpcWalletGateway struct {
	client pb.WalletServiceClient
}

func NewGrpcWalletGateway(conn *grpc.ClientConn) *GrpcWalletGateway {
	client := pb.NewWalletServiceClient(conn)
	return &GrpcWalletGateway{client: client}
}

func (g *GrpcWalletGateway) GetBalance(ctx context.Context, userID string) (*domain.ExternalWallet, error) {
	response, err := g.client.GetBalance(ctx, &pb.GetUserBalanceRequest{UserId: userID})
	if err != nil {
		return nil, err
	}
	return MapProtoToWallet(response), nil
}
