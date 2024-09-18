package storage

import (
	"context"
	"tg_go_coins_service/genproto/coins_service"
)

type StorageI interface {
	CloseDB()
	Coin() CoinRepoI
	GetBuyOrSell() GetBuyOrSellRepoI
	History() HistoryUserRepoI
	TelegramPremium() TelegramPremiumRepoI
	NFT() NFTRepoI
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

type HistoryUserRepoI interface {
	HistoryUser(ctx context.Context, req *coins_service.HistoryUserRequest) (resp *coins_service.HistoryUserResponse, err error)
	HistoryUserAll(ctx context.Context) (resp *coins_service.HistoryUserResponse, err error)
}

type TelegramPremiumRepoI interface {
	CreatePrice(ctx context.Context, req *coins_service.CreateTelegramPremiumPrice) (resp *coins_service.TelegramPremiumPrice, err error)
	CreatePremium(ctx context.Context, req *coins_service.CreateTelegramPremium) (resp *coins_service.TelegramPremium, err error)
	GetPremiumById(ctx context.Context, req *coins_service.TelegramPriemiumPrimaryKey) (resp *coins_service.TelegramPremium, err error)
	UpdateTransactionStatus(ctx context.Context, req *coins_service.UpdateStatus) (int64, error)
	PremiumTransaction(ctx context.Context, req *coins_service.PremiumTransactionRequest) error
	GetList(ctx context.Context, req *coins_service.GetListPremiumRequest) (resp *coins_service.GetPremiumTransactionResponse, err error)
	GetPremiumList(ctx context.Context, req *coins_service.GetPremiumListRequest) (resp *coins_service.GetPremiumListResponse, err error)
	GetPremiumTransactionById(ctx context.Context, req *coins_service.GetPremiumTransactionPrimaryKey) (resp *coins_service.GetPremiumTransactionId, err error)
}

type NFTRepoI interface {
	Create(ctx context.Context, req *coins_service.CreateNFT) (resp *coins_service.NFT, err error)
	GetById(ctx context.Context, req *coins_service.NFTPrimaryKey) (resp *coins_service.NFT, err error)
	GetAll(ctx context.Context, req *coins_service.GetListNFTRequest) (resp *coins_service.GetListNFTResponse, err error)
	Update(ctx context.Context, req *coins_service.UpdateNFT) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *coins_service.NFTPrimaryKey) error
}
