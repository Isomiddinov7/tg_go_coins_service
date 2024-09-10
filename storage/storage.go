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
}
