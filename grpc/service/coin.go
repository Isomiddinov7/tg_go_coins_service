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

type CoinService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*coins_service.UnimplementedCoinsServiceServer
}

func NewCoinService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CoinService {
	return &CoinService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CoinService) Create(ctx context.Context, req *coins_service.CreateCoin) (resp *coins_service.CoinPrimaryKey, err error) {

	i.log.Info("---CreateCoin------>", logger.Any("req", req))
	resp, err = i.strg.Coin().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateCoin->Coin->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *CoinService) GetById(ctx context.Context, req *coins_service.CoinPrimaryKey) (resp *coins_service.Coin, err error) {
	i.log.Info("---GetCoinByID------>", logger.Any("req", req))

	resp, err = i.strg.Coin().GetByID(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCoinByID->Coin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *CoinService) GetList(ctx context.Context, req *coins_service.GetListCoinRequest) (resp *coins_service.GetListCoinResponse, err error) {

	i.log.Info("---GetCoins------>", logger.Any("req", req))

	resp, err = i.strg.Coin().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCoins->Coin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CoinService) Update(ctx context.Context, req *coins_service.UpdateCoin) (resp *coins_service.Coin, err error) {

	i.log.Info("---UpdateCoin------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Coin().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateCoin--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Coin().GetByID(ctx, &coins_service.CoinPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CoinService) Delete(ctx context.Context, req *coins_service.CoinPrimaryKey) (resp *coins_service.Empty, err error) {

	i.log.Info("---DeleteCoin------>", logger.Any("req", req))

	err = i.strg.Coin().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteCoin->Coin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &coins_service.Empty{}, nil
}
