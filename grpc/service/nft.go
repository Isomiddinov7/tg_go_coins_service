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

type NFTService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*coins_service.UnimplementedNFTServiceServer
}

func NewNFTService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *NFTService {
	return &NFTService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *NFTService) Create(ctx context.Context, req *coins_service.CreateNFT) (resp *coins_service.NFT, err error) {
	i.log.Info("---CreateNFT------>", logger.Any("req", req))
	resp, err = i.strg.NFT().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateNFT->NFT->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *NFTService) GetById(ctx context.Context, req *coins_service.NFTPrimaryKey) (resp *coins_service.NFT, err error) {
	i.log.Info("---GetNFTByID------>", logger.Any("req", req))
	resp, err = i.strg.NFT().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetNFTByID->NFT->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *NFTService) GetAll(ctx context.Context, req *coins_service.GetListNFTRequest) (resp *coins_service.GetListNFTResponse, err error) {
	i.log.Info("---GetAllNFT------>", logger.Any("req", req))
	resp, err = i.strg.NFT().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAllNFT->NFT->GetAll--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *NFTService) Update(ctx context.Context, req *coins_service.UpdateNFT) (resp *coins_service.NFT, err error) {
	i.log.Info("---UpdateNFT------>", logger.Any("req", req))
	rowsAffected, err := i.strg.NFT().Update(ctx, req)
	if err != nil {
		i.log.Error("!!!UpdateNFT--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}

func (i *NFTService) Delete(ctx context.Context, req *coins_service.NFTPrimaryKey) (resp *coins_service.Empty, err error) {

	i.log.Info("---DeleteNFT------>", logger.Any("req", req))

	err = i.strg.NFT().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteNFT->NFT->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &coins_service.Empty{}, nil
}
