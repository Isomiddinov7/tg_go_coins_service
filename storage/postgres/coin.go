package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"tg_go_coins_service/genproto/coins_service"
	"tg_go_coins_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type coinRepo struct {
	db *pgxpool.Pool
}

func NewCoinRepo(db *pgxpool.Pool) storage.CoinRepoI {
	return &coinRepo{
		db: db,
	}
}

func (r *coinRepo) Create(ctx context.Context, req *coins_service.CreateCoin) (resp *coins_service.CoinPrimaryKey, err error) {
	var (
		id    = uuid.NewString()
		query = `
			INSERT INTO "coins"(
				"id",
				"name",
				"coin_buy_price",
				"coin_sell_price",
				"address",
				"card_number",
				"image",
				"status"
			) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
	)
	_, err = r.db.Exec(ctx,
		query,
		id,
		req.Name,
		req.CoinBuyPrice,
		req.CoinSellPrice,
		req.Address,
		req.CardNumber,
		req.ImageId,
		req.Status,
	)
	if err != nil {
		return nil, err
	}

	if len(req.Halfcoins) > 0 {
		var (
			queryHalf = `
			INSERT INTO "half_coins_price"(
				"coin_id",
				"halfCoinAmount",
				"halfCoinPrice"
			) VALUES ($1, $2, $3)
		`

			halfCoin = &coins_service.HalfCoinPrice{}
		)

		for i := range req.Halfcoins {
			halfCoin.HalfCoinAmount = req.Halfcoins[i].HalfCoinAmount
			halfCoin.HalfCoinPrice = req.Halfcoins[i].HalfCoinPrice

			_, err = r.db.Exec(ctx, queryHalf,
				id,
				halfCoin.HalfCoinAmount,
				halfCoin.HalfCoinPrice,
			)
			if err != nil {
				return nil, err
			}
		}

	}

	return &coins_service.CoinPrimaryKey{
		Id: id,
	}, nil
}

func (r *coinRepo) GetByID(ctx context.Context, req *coins_service.CoinPrimaryKey) (*coins_service.Coin, error) {
	queryCoin := `
		SELECT
			"id",
			"name",
			"coin_buy_price",
			"coin_sell_price",
			"address",
			"card_number",
			"status",
			"image",
			"created_at",
			"updated_at"
		FROM "coins"
		WHERE id = $1
	`

	var (
		id              sql.NullString
		name            sql.NullString
		coin_buy_price  sql.NullString
		coin_sell_price sql.NullString
		address         sql.NullString
		image           sql.NullString
		card_number     sql.NullString
		status          sql.NullString
		created_at      sql.NullString
		updated_at      sql.NullString
	)

	err := r.db.QueryRow(ctx, queryCoin, req.Id).Scan(
		&id,
		&name,
		&coin_buy_price,
		&coin_sell_price,
		&address,
		&card_number,
		&status,
		&image,
		&created_at,
		&updated_at,
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

	rows, err := r.db.Query(ctx, queryHalf, req.Id)
	if err != nil {
		return nil, err
	}
	halfPrices := []*coins_service.HalfCoinPrice{}
	for rows.Next() {
		var (
			halfPrice      = &coins_service.HalfCoinPrice{}
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
		halfPrice = &coins_service.HalfCoinPrice{
			HalfCoinAmount: halfCoinAmount.String,
			HalfCoinPrice:  halfCoinPrice.String,
		}
		halfPrices = append(halfPrices, halfPrice)
	}
	return &coins_service.Coin{
		Id:            id.String,
		Name:          name.String,
		CoinBuyPrice:  coin_buy_price.String,
		CoinSellPrice: coin_sell_price.String,
		Halfcoins:     halfPrices,
		Address:       address.String,
		CardNumber:    card_number.String,
		Status:        status.String,
		ImageId:       image.String,
		CreatedAt:     created_at.String,
		UpdatedAt:     updated_at.String,
	}, nil

}

func (r *coinRepo) GetAll(ctx context.Context, req *coins_service.GetListCoinRequest) (*coins_service.GetListCoinResponse, error) {
	var (
		resp   coins_service.GetListCoinResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.Search) > 0 {
		where += " AND name ILIKE" + " '%" + req.Search + "%'"
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			"id",
			"name",
			"coin_buy_price",
			"coin_sell_price",
			"address",
			"card_number",
			"status",
			"image",
			"created_at",
			"updated_at"
		FROM "coins"
		`

	query += where + sort + offset + limit

	rowsCoins, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rowsCoins.Close()

	for rowsCoins.Next() {
		var (
			coin            coins_service.Coin
			id              sql.NullString
			name            sql.NullString
			coin_buy_price  sql.NullString
			coin_sell_price sql.NullString
			address         sql.NullString
			image           sql.NullString
			card_number     sql.NullString
			status          sql.NullString
			created_at      sql.NullString
			updated_at      sql.NullString
		)

		err = rowsCoins.Scan(
			&resp.Count,
			&id,
			&name,
			&coin_buy_price,
			&coin_sell_price,
			&address,
			&card_number,
			&status,
			&image,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}

		queryHalf := `
			SELECT
				"halfCoinAmount",
				"halfCoinPrice"
			FROM "half_coins_price"
			WHERE "coin_id" = $1
		`

		rowsHalf, err := r.db.Query(ctx, queryHalf, id.String)
		if err != nil {
			return nil, err
		}
		defer rowsHalf.Close()

		halfPrices := []*coins_service.HalfCoinPrice{}
		for rowsHalf.Next() {
			var (
				halfPrice      = &coins_service.HalfCoinPrice{}
				halfCoinAmount sql.NullString
				halfCoinPrice  sql.NullString
			)

			err = rowsHalf.Scan(
				&halfCoinAmount,
				&halfCoinPrice,
			)
			if err != nil {
				return nil, err
			}
			halfPrice = &coins_service.HalfCoinPrice{
				HalfCoinAmount: halfCoinAmount.String,
				HalfCoinPrice:  halfCoinPrice.String,
			}
			halfPrices = append(halfPrices, halfPrice)
		}

		coin = coins_service.Coin{
			Id:            id.String,
			Name:          name.String,
			CoinBuyPrice:  coin_buy_price.String,
			CoinSellPrice: coin_sell_price.String,
			Halfcoins:     halfPrices,
			Address:       address.String,
			CardNumber:    card_number.String,
			Status:        status.String,
			ImageId:       image.String,
			CreatedAt:     created_at.String,
			UpdatedAt:     updated_at.String,
		}

		resp.Coins = append(resp.Coins, &coin)
	}
	return &resp, nil
}

func (r *coinRepo) Update(ctx context.Context, req *coins_service.UpdateCoin) (int64, error) {
	var (
		query = `
		UPDATE "coins"
			SET
				"name" = $2,
				"coin_buy_price" = $3,
				"coin_sell_price" = $4,
				"address" = $5,
				"card_number" = $6,
				"status" = $7,
				"updated_at" = NOW()
		WHERE "id" = $1`
	)

	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Name,
		req.CoinBuyPrice,
		req.CoinSellPrice,
		req.Address,
		req.CardNumber,
		req.Status,
	)
	if err != nil {
		return 0, err
	}

	_, err = r.db.Exec(ctx, `DELETE FROM "half_coins_price" WHERE "coin_id" = $1`, req.Id)
	if err != nil {
		return 0, err
	}

	if len(req.Halfcoins) > 0 {
		var (
			queryHalf = `
			INSERT INTO "half_coins_price"(
				"coin_id",
				"halfCoinAmount",
				"halfCoinPrice"
			) VALUES ($1, $2, $3)
		`

			halfCoin = &coins_service.HalfCoinPrice{}
		)

		for i := range req.Halfcoins {
			halfCoin.HalfCoinAmount = req.Halfcoins[i].HalfCoinAmount
			halfCoin.HalfCoinPrice = req.Halfcoins[i].HalfCoinPrice

			_, err = r.db.Exec(ctx, queryHalf,
				req.Id,
				halfCoin.HalfCoinAmount,
				halfCoin.HalfCoinPrice,
			)
			if err != nil {
				return 0, err
			}
		}

	}

	return rowsAffected.RowsAffected(), nil
}

func (r *coinRepo) Delete(ctx context.Context, req *coins_service.CoinPrimaryKey) error {
	_, _ = r.db.Exec(ctx, `DELETE FROM "half_coins_price" WHERE "coin_id" = $1`, req.Id)
	_, err := r.db.Exec(ctx, `DELETE FROM "coins" WHERE id = $1`, req.Id)
	return err
}
