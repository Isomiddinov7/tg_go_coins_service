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

type PremiumService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	coins_service.UnimplementedTelegramPremiumServiceServer
}

func NewTelegramPremiumService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *PremiumService {
	return &PremiumService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *PremiumService) CreatePrice(ctx context.Context, req *coins_service.CreateTelegramPremiumPrice) (resp *coins_service.TelegramPremiumPrice, err error) {

	i.log.Info("---CreatePrice------>", logger.Any("req", req))
	resp, err = i.strg.TelegramPremium().CreatePrice(ctx, req)
	if err != nil {
		i.log.Error("!!!CreatePrice->TelegramPremium->CreatePrice--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *PremiumService) CreatePremium(ctx context.Context, req *coins_service.CreateTelegramPremium) (resp *coins_service.TelegramPremium, err error) {

	i.log.Info("---CreatePremium------>", logger.Any("req", req))
	resp, err = i.strg.TelegramPremium().CreatePremium(ctx, req)
	if err != nil {
		i.log.Error("!!!CreatePremium->TelegramPremium->CreatePremium--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *PremiumService) GetPremiumById(ctx context.Context, req *coins_service.TelegramPriemiumPrimaryKey) (resp *coins_service.TelegramPremium, err error) {

	i.log.Info("---GetPremiumById------>", logger.Any("req", req))
	resp, err = i.strg.TelegramPremium().GetPremiumById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetPremiumById->TelegramPremium->GetPremiumById--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *PremiumService) UpdateTransactionStatus(ctx context.Context, req *coins_service.UpdateStatus) (resp *coins_service.Empty, err error) {

	i.log.Info("---UpdateTransactionStatus------>", logger.Any("req", req))
	_, err = i.strg.TelegramPremium().UpdateTransactionStatus(ctx, req)
	if err != nil {
		i.log.Error("!!!UpdateTransactionStatus->TelegramPremium->UpdateTransactionStatus--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *PremiumService) PremiumTransaction(ctx context.Context, req *coins_service.PremiumTransactionRequest) (resp *coins_service.Empty, err error) {

	i.log.Info("---PremiumTransaction------>", logger.Any("req", req))
	err = i.strg.TelegramPremium().PremiumTransaction(ctx, req)
	if err != nil {
		i.log.Error("!!!PremiumTransaction->TelegramPremium->PremiumTransaction--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *PremiumService) GetList(ctx context.Context, req *coins_service.GetListPremiumRequest) (resp *coins_service.GetPremiumTransactionResponse, err error) {

	i.log.Info("---GetList------>", logger.Any("req", req))
	resp, err = i.strg.TelegramPremium().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetList->TelegramPremium->GetList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *PremiumService) GetPremiumList(ctx context.Context, req *coins_service.GetPremiumListRequest) (resp *coins_service.GetPremiumListResponse, err error) {

	i.log.Info("---GetPremiumList------>", logger.Any("req", req))
	resp, err = i.strg.TelegramPremium().GetPremiumList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetPremiumList->TelegramPremium->GetList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *PremiumService) GetPremiumTransactionById(ctx context.Context, req *coins_service.GetPremiumTransactionPrimaryKey) (resp *coins_service.GetPremiumTransactionId, err error) {
	i.log.Info("---GetPremiumTransactionById------>", logger.Any("req", req))
	resp, err = i.strg.TelegramPremium().GetPremiumTransactionById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetPremiumTransactionById->TelegramPremium->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}
