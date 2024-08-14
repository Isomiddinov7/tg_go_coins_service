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

type ImageService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*coins_service.UnimplementedImagesServiceServer
}

func NewImageService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *ImageService {
	return &ImageService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (i *ImageService) ImageUpload(ctx context.Context, req *coins_service.ImageData) (*coins_service.ImagePrimaryKey, error) {
	i.log.Info("---FileUpload------>", logger.Any("req", req))
	resp, err := i.strg.FileImage().FileUpload(ctx, &coins_service.ImageData{Id: req.Id, ImageLink: req.ImageLink})
	if err != nil {
		i.log.Error("!!!FileUpload--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (i *ImageService) ImageDelete(ctx context.Context, req *coins_service.ImagePrimaryKey) (resp *coins_service.Empty, err error) {
	i.log.Info("---FileDelete------>", logger.Any("req", req))
	_, err = i.strg.FileImage().FileDelete(ctx, &coins_service.ImagePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!FileDelete--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, nil
}
