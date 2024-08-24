package service

import (
	"context"
	"tg_go_coins_service/config"
	"tg_go_coins_service/genproto/coins_service"
	"tg_go_coins_service/grpc/client"
	"tg_go_coins_service/pkg/logger"
	"tg_go_coins_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HistoryService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*coins_service.UnimplementedHistoryServiceServer
}

func NewHistoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *HistoryService {
	return &HistoryService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CoinService) UserHistory(ctx context.Context, req *coins_service.HistoryUserRequest) (resp *coins_service.HistoryUserResponse, err error) {

	i.log.Info("---UserHistory------>", logger.Any("req", req))

	resp, err = i.strg.History().HistoryUser(ctx, req)
	if err != nil {
		i.log.Error("!!!UserHistory->Coin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}
