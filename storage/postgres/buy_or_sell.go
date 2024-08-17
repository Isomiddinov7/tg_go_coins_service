package postgres

import (
	"context"
	"database/sql"
	"tg_go_coins_service/genproto/coins_service"
	"tg_go_coins_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/cast"
)

type buysellRepo struct {
	db *pgxpool.Pool
}

func NewBuyOrSellRepo(db *pgxpool.Pool) storage.GetBuyOrSellRepoI {
	return &buysellRepo{
		db: db,
	}
}

func (r *buysellRepo) GetSell(ctx context.Context, req *coins_service.BuyOrSellRequest) (*coins_service.BuyOrSellResponse, error) {
	var (
		queryAdmin = `
				SELECT
					"address"
				FROM "coins"
				WHERE "id" = $1
			`
		admin_address sql.NullString
	)

	err := r.db.QueryRow(ctx, queryAdmin, req.CoinId).Scan(
		&admin_address,
	)
	if err != nil {
		return nil, err
	}

	var (
		queryHalf = `
			SELECT
				"halfCoinAmount",
				"halfCoinPrice"
			FROM "half_coins_price"
			WHERE "coin_id" = $1
		`
	)

	rows, err := r.db.Query(ctx, queryHalf, req.CoinId)
	if err != nil {
		return nil, err
	}
	halfPrices := []*coins_service.HalfCoinPrice{}
	for rows.Next() {
		var (
			halfPrice      = coins_service.HalfCoinPrice{}
			halfCoinAmount sql.NullString
			halfCoinPrice  sql.NullString
		)

		err = rows.Scan(
			&halfCoinAmount,
			&halfCoinPrice,
		)
		if err != nil {
			return nil, err
		}
		halfPrice = coins_service.HalfCoinPrice{
			HalfCoinAmount: halfCoinAmount.String,
			HalfCoinPrice:  halfCoinPrice.String,
		}
		halfPrices = append(halfPrices, &halfPrice)
	}
	var (
		coin_sell_price sql.NullString
		summ            float64
	)
	if cast.ToInt64(req.CoinAmount) >= 1 {
		query1 := `
			SELECT
				"coin_sell_price"
			FROM "coins"
			WHERE id = $1
		`
		err = r.db.QueryRow(ctx, query1, req.CoinId).Scan(
			&coin_sell_price,
		)
		if err != nil {
			return nil, err
		}
		summ = cast.ToFloat64(req.CoinAmount) * cast.ToFloat64(coin_sell_price.String)

	} else {
		for i := range halfPrices {
			if cast.ToFloat64(req.CoinAmount) == cast.ToFloat64(halfPrices[i].HalfCoinAmount) {
				coin_sell_price.String = cast.ToString(halfPrices[i].HalfCoinPrice)
			}
		}
		summ = cast.ToFloat64(coin_sell_price.String)
	}

	return &coins_service.BuyOrSellResponse{
		AdminAddress: admin_address.String,
		CoinAmount:   req.CoinAmount,
		AmountPrice:  cast.ToString(summ),
	}, nil
}

func (r *buysellRepo) GetBuy(ctx context.Context, req *coins_service.BuyOrSellRequest) (*coins_service.BuyOrSellResponse, error) {
	var (
		queryAdmin = `
				SELECT
					"card_number"
				FROM "coins"
				WHERE "id" = $1
			`
		admin_address sql.NullString
	)

	err := r.db.QueryRow(ctx, queryAdmin, req.CoinId).Scan(
		&admin_address,
	)
	if err != nil {
		return nil, err
	}
	var (
		queryHalf = `
			SELECT
				"halfCoinAmount",
				"halfCoinPrice"
			FROM "half_coins_price"
			WHERE "coin_id" = $1
		`
	)

	rows, err := r.db.Query(ctx, queryHalf, req.CoinId)
	if err != nil {
		return nil, err
	}
	halfPrices := []*coins_service.HalfCoinPrice{}
	for rows.Next() {
		var (
			halfPrice      = coins_service.HalfCoinPrice{}
			halfCoinAmount sql.NullString
			halfCoinPrice  sql.NullString
		)

		err = rows.Scan(
			&halfCoinAmount,
			&halfCoinPrice,
		)
		if err != nil {
			return nil, err
		}
		halfPrice = coins_service.HalfCoinPrice{
			HalfCoinAmount: halfCoinAmount.String,
			HalfCoinPrice:  halfCoinPrice.String,
		}
		halfPrices = append(halfPrices, &halfPrice)
	}

	var (
		coin_buy_price sql.NullString
		summ           float64
	)

	if cast.ToInt64(req.CoinAmount) >= 1 {
		query1 := `
			SELECT
				"coin_buy_price"
			FROM "coins"
			WHERE id = $1
		`
		err = r.db.QueryRow(ctx, query1, req.CoinId).Scan(
			&coin_buy_price,
		)
		if err != nil {
			return nil, err
		}
		summ = cast.ToFloat64(req.CoinAmount) * cast.ToFloat64(coin_buy_price.String)

	} else {
		for i := range halfPrices {
			if cast.ToFloat64(req.CoinAmount) == cast.ToFloat64(halfPrices[i].HalfCoinAmount) {
				coin_buy_price.String = cast.ToString(halfPrices[i].HalfCoinPrice)
			}
		}
		summ = cast.ToFloat64(coin_buy_price.String)
	}

	return &coins_service.BuyOrSellResponse{
		AdminAddress: admin_address.String,
		CoinAmount:   req.CoinAmount,
		AmountPrice:  cast.ToString(summ),
	}, nil
}
