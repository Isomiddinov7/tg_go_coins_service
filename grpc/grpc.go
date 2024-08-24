package grpc

import (
	"tg_go_coins_service/config"
	"tg_go_coins_service/genproto/coins_service"
	"tg_go_coins_service/grpc/client"
	"tg_go_coins_service/grpc/service"
	"tg_go_coins_service/pkg/logger"
	"tg_go_coins_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()
	coins_service.RegisterCoinsServiceServer(grpcServer, service.NewCoinService(cfg, log, strg, srvc))
	coins_service.RegisterBuyOrSellServer(grpcServer, service.NewBuyOrSellService(cfg, log, strg, srvc))
	coins_service.RegisterImagesServiceServer(grpcServer, service.NewImageService(cfg, log, strg, srvc))
	coins_service.RegisterHistoryServiceServer(grpcServer, service.NewHistoryService(cfg, log, strg, srvc))
	reflection.Register(grpcServer)
	return
}
