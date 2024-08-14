package postgres

import (
	"context"
	"fmt"
	"tg_go_coins_service/config"
	"tg_go_coins_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	coin     storage.CoinRepoI
	buy_sell storage.GetBuyOrSellRepoI
	image    storage.ImagesRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Coin() storage.CoinRepoI {
	if s.coin == nil {
		s.coin = NewCoinRepo(s.db)
	}

	return s.coin
}

func (s *Store) GetBuyOrSell() storage.GetBuyOrSellRepoI {
	if s.buy_sell == nil {
		s.buy_sell = NewBuyOrSellRepo(s.db)
	}

	return s.buy_sell
}

func (s *Store) FileImage() storage.ImagesRepoI {
	if s.image == nil {
		s.image = NewFileImageRepo(s.db)
	}

	return s.image
}
