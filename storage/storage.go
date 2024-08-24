package storage

import (
	"context"
	"tg_go_coins_service/genproto/coins_service"
)

type StorageI interface {
	CloseDB()
	Coin() CoinRepoI
	GetBuyOrSell() GetBuyOrSellRepoI
	FileImage() ImagesRepoI
	History() HistoryUserRepoI
}

type CoinRepoI interface {
	Create(ctx context.Context, req *coins_service.CreateCoin) (resp *coins_service.CoinPrimaryKey, err error)
	GetByID(ctx context.Context, req *coins_service.CoinPrimaryKey) (resp *coins_service.Coin, err error)
	GetAll(ctx context.Context, req *coins_service.GetListCoinRequest) (resp *coins_service.GetListCoinResponse, err error)
	Update(ctx context.Context, req *coins_service.UpdateCoin) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *coins_service.CoinPrimaryKey) error
}

type GetBuyOrSellRepoI interface {
	GetSell(ctx context.Context, req *coins_service.BuyOrSellRequest) (*coins_service.BuyOrSellResponse, error)
	GetBuy(ctx context.Context, req *coins_service.BuyOrSellRequest) (*coins_service.BuyOrSellResponse, error)
}

type ImagesRepoI interface {
	FileUpload(ctx context.Context, req *coins_service.ImageData) (resp *coins_service.ImagePrimaryKey, err error)
	FileDelete(ctx context.Context, req *coins_service.ImagePrimaryKey) (resp *coins_service.Empty, err error)
}

type HistoryUserRepoI interface {
	HistoryUser(ctx context.Context, req *coins_service.HistoryUserRequest) (resp *coins_service.HistoryUserResponse, err error)
}
