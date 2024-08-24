package postgres

import (
	"context"
	"database/sql"
	"tg_go_coins_service/genproto/coins_service"
	"tg_go_coins_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type historyRepo struct {
	db *pgxpool.Pool
}

func NewHistoryRepo(db *pgxpool.Pool) storage.HistoryUserRepoI {
	return &historyRepo{
		db: db,
	}
}

func (r *historyRepo) HistoryUser(ctx context.Context, req *coins_service.HistoryUserRequest) (*coins_service.HistoryUserResponse, error) {
	var (
		resp  coins_service.HistoryUserResponse
		query = `
			SELECT 
				id,
				status,
				user_confirmation_img,
				coin_amount,
				coin_price,
				all_price,
				user_address,
				payment_card
				created_at
			FROM "user_transaction"
			WHERE user_id = $1
		`
	)

	rows, err := r.db.Query(ctx, query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var (
			history               coins_service.HistoriesUser
			id                    sql.NullString
			status                sql.NullString
			user_confirmation_img sql.NullString
			coin_amount           sql.NullString
			coin_price            sql.NullString
			all_price             sql.NullString
			user_address          sql.NullString
			payment_card          sql.NullString
			created_at            sql.NullString
		)

		err = rows.Scan(
			&id,
			&status,
			&user_confirmation_img,
			&coin_amount,
			&coin_price,
			&all_price,
			&user_address,
			&payment_card,
			&created_at,
		)
		if err != nil {
			return nil, err
		}

		history = coins_service.HistoriesUser{
			Id:         id.String,
			Status:     status.String,
			CoinImg:    user_confirmation_img.String,
			CoinAmount: coin_amount.String,
			CoinPrice:  coin_price.String,
			AllPrice:   all_price.String,
			Address:    user_address.String,
			CardNumber: payment_card.String,
			DateTime:   created_at.String,
		}

		resp.Histories = append(resp.Histories, &history)
	}
	return &resp, nil
}
