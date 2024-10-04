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

type CoinNftService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	coins_service.UnimplementedCoinNFTServiceServer
}

func NewCoinNftService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CoinNftService {
	return &CoinNftService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CoinNftService) Create(ctx context.Context, req *coins_service.CoinNFTCreate) (resp *coins_service.CoinNFT, err error) {

	i.log.Info("---CreateCoin------>", logger.Any("req", req))
	resp, err = i.strg.CoinNFT().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateCoin->Coin->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *CoinNftService) GetById(ctx context.Context, req *coins_service.CoinNFTPrimaryKey) (resp *coins_service.CoinNFT, err error) {
	i.log.Info("---GetCoinByID------>", logger.Any("req", req))

	resp, err = i.strg.CoinNFT().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCoinByID->Coin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *CoinNftService) GetList(ctx context.Context, req *coins_service.GetListCoinNFTRequest) (resp *coins_service.GetListCoinNFTResponse, err error) {

	i.log.Info("---GetCoins------>", logger.Any("req", req))

	resp, err = i.strg.CoinNFT().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCoins->Coin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *CoinNftService) Update(ctx context.Context, req *coins_service.CoinNFTUpdate) (resp *coins_service.CoinNFT, err error) {

	i.log.Info("---UpdateCoin------>", logger.Any("req", req))

	rowsAffected, err := i.strg.CoinNFT().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateCoin--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.CoinNFT().GetById(ctx, &coins_service.CoinNFTPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CoinNftService) Delete(ctx context.Context, req *coins_service.CoinNFTPrimaryKey) (resp *coins_service.Empty, err error) {

	i.log.Info("---DeleteCoin------>", logger.Any("req", req))

	err = i.strg.CoinNFT().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteCoin->Coin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &coins_service.Empty{}, nil
}
