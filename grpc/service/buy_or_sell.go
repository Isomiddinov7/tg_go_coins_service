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

type BuyOrSellService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	coins_service.UnimplementedBuyOrSellServer
}

func NewBuyOrSellService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BuyOrSellService {
	return &BuyOrSellService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *BuyOrSellService) GetSell(ctx context.Context, req *coins_service.BuyOrSellRequest) (resp *coins_service.BuyOrSellResponse, err error) {

	i.log.Info("---GetCoinSell------>", logger.Any("req", req))

	resp, err = i.strg.GetBuyOrSell().GetSell(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCoins Sell->Coin Sell->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BuyOrSellService) GetBuy(ctx context.Context, req *coins_service.BuyOrSellRequest) (resp *coins_service.BuyOrSellResponse, err error) {

	i.log.Info("---GetCoinBuy------>", logger.Any("req", req))

	resp, err = i.strg.GetBuyOrSell().GetBuy(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCoins Buy->Coin Buy->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}
