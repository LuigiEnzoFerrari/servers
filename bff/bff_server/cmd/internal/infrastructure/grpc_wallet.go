package infrastructure

import (
	"context"
	"google.golang.org/grpc"
	pb "github.com/LuigiEnzoFerrari/servers/bff/bff_server/api/proto/wallet/v1"
)

type GrpcWalletGateway struct {
	client pb.WalletServiceClient
}

func NewGrpcWalletGateway(conn *grpc.ClientConn) *GrpcWalletGateway {
	client := pb.NewWalletServiceClient(conn)
	return &GrpcWalletGateway{client: client}
}

func (g *GrpcWalletGateway) GetBalance(ctx context.Context, userID string) (*GetUserBalanceResponse, error) {
	response, err := g.client.GetBalance(ctx, &pb.GetUserBalanceRequest{UserId: userID})
	if err != nil {
		return nil, err
	}
	return MapProtoToWallet(response), nil
}

