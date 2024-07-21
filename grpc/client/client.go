package client

import (
	"tg_go_coins_service/config"
	"tg_go_coins_service/genproto/coins_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	CoinService() coins_service.CoinsServiceClient
}

type grpcClients struct {
	coinService coins_service.CoinsServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connCoinsService, err := grpc.NewClient(
		cfg.CoinsServiceHost+cfg.CoinsGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &grpcClients{
		coinService: coins_service.NewCoinsServiceClient(connCoinsService),
	}, nil
}

func (g *grpcClients) CoinService() coins_service.CoinsServiceClient {
	return g.coinService
}
