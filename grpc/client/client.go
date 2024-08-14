package client

import (
	"tg_go_coins_service/config"
	"tg_go_coins_service/genproto/coins_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	CoinService() coins_service.CoinsServiceClient
	BuyOrSellService() coins_service.BuyOrSellClient
	FileImageService() coins_service.ImagesServiceClient
}

type grpcClients struct {
	coinService      coins_service.CoinsServiceClient
	buyorsellService coins_service.BuyOrSellClient
	fileimageService coins_service.ImagesServiceClient
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
		coinService:      coins_service.NewCoinsServiceClient(connCoinsService),
		buyorsellService: coins_service.NewBuyOrSellClient(connCoinsService),
		fileimageService: coins_service.NewImagesServiceClient(connCoinsService),
	}, nil
}

func (g *grpcClients) CoinService() coins_service.CoinsServiceClient {
	return g.coinService
}

func (g *grpcClients) BuyOrSellService() coins_service.BuyOrSellClient {
	return g.buyorsellService
}

func (g *grpcClients) FileImageService() coins_service.ImagesServiceClient {
	return g.fileimageService
}
