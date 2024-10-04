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
				ut.id,
				c.name,
				ut.status,
				ut.user_confirmation_img,
				ut.coin_amount,
				ut.coin_price,
				ut.all_price,
				ut.user_address,
				ut.payment_card,
				ut.created_at
			FROM "user_transaction" as ut
			JOIN "coins" as c ON c.id = ut.coin_id
			WHERE ut.user_id = $1
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
			name                  sql.NullString
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
			&name,
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
			Name:       name.String,
			Status:     status.String,
			ConfirmImg: user_confirmation_img.String,
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

func (r *historyRepo) HistoryUserAll(ctx context.Context) (*coins_service.HistoryUserResponse, error) {
	var (
		resp  coins_service.HistoryUserResponse
		query = `
			SELECT 
				ut.id,
				c.name
				ut.status,
				ut.user_confirmation_img,
				ut.coin_amount,
				ut.coin_price,
				ut.all_price,
				ut.user_address,
				ut.payment_card,
				ut.created_at
			FROM "user_transaction"
			JOIN "coins" as c ON c.id = ut.coin_id
		`
	)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var (
			history               coins_service.HistoriesUser
			id                    sql.NullString
			name                  sql.NullString
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
			&name,
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
			Name:       name.String,
			Status:     status.String,
			ConfirmImg: user_confirmation_img.String,
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

func (r *historyRepo) HistoryMessage(ctx context.Context, req *coins_service.HistoryUserRequest) (resp *coins_service.HistoryMessageResponse, err error) {
	var (
		query = `
			SELECT
				ut."id",
				c."name",
				ut."coin_id",
				ut."user_id",
				ut."user_confirmation_img",
				ut."coin_price",
				ut."coin_amount",
				ut."all_price",
				ut."status",
				ut."user_address",
				ut."card_name",
				ut."payment_card",
				ut."message",
				ut."transaction_status",
				ut."created_at"
			FROM "user_transaction" as ut
			JOIN "coins" as c ON c.id = ut.coin_id
			WHERE ut."user_id" = $1
		`

		id                    sql.NullString
		name                  sql.NullString
		coin_id               sql.NullString
		user_id               sql.NullString
		user_confirmation_img sql.NullString
		coin_price            sql.NullString
		coin_amount           sql.NullString
		all_price             sql.NullString
		status                sql.NullString
		user_address          sql.NullString
		card_name             sql.NullString
		payment_card          sql.NullString
		message               sql.NullString
		transaction_status    sql.NullString
		created_at            sql.NullString
		data                  []*coins_service.HistoryUserWithStatus
	)

	rows, err := r.db.Query(ctx, query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&id,
			&name,
			&coin_id,
			&user_id,
			&user_confirmation_img,
			&coin_price,
			&coin_amount,
			&all_price,
			&status,
			&user_address,
			&card_name,
			&payment_card,
			&message,
			&transaction_status,
			&created_at,
		)
		if err != nil {
			return nil, err
		}

		var (
			queryMessage = `
				SELECT
					pm."message",
					pm."read",
					pm."file"
				FROM pay_message as pm
				WHERE pm.user_id = $1 AND pm.user_transaction_id = $2
			`
			text sql.NullString
			read sql.NullString
			file sql.NullString
		)

		err = r.db.QueryRow(ctx, queryMessage, req.UserId, id.String).Scan(&text, &read, &file)
		if err == sql.ErrNoRows {
			result := &coins_service.HistoryUserWithStatus{
				HistoryUser: &coins_service.HistoriesUser{
					Id:                id.String,
					Name:              name.String,
					Status:            status.String,
					ConfirmImg:        user_confirmation_img.String,
					CoinAmount:        coin_amount.String,
					CoinPrice:         coin_price.String,
					AllPrice:          all_price.String,
					Address:           user_address.String,
					CardNumber:        payment_card.String,
					DateTime:          created_at.String,
					TransactionStatus: transaction_status.String,
					CoinId:            coin_id.String,
					UserId:            user_id.String,
				},
				HistoryStatus: nil,
			}
			data = append(data, result)
			continue
		} else if err != nil {
			return nil, err
		}

		result := &coins_service.HistoryUserWithStatus{
			HistoryUser: &coins_service.HistoriesUser{
				Id:                id.String,
				Name:              name.String,
				Status:            status.String,
				ConfirmImg:        user_confirmation_img.String,
				CoinAmount:        coin_amount.String,
				CoinPrice:         coin_price.String,
				AllPrice:          all_price.String,
				Address:           user_address.String,
				CardNumber:        payment_card.String,
				DateTime:          created_at.String,
				TransactionStatus: transaction_status.String,
				CoinId:            coin_id.String,
				UserId:            user_id.String,
			},
			HistoryStatus: &coins_service.TransactionStatus{
				Text:    text.String,
				Status:  read.String,
				Message: file.String,
			},
		}

		data = append(data, result)
	}

	resp = &coins_service.HistoryMessageResponse{
		HistoryWithStatus: data,
	}
	return resp, nil
}

func (r *historyRepo) UpdateHistoryRead(ctx context.Context, req *coins_service.HistoryUserRequest) (rowsAffected int64, err error) {
	var (
		query = `
			UPDATE "pay_message"
				SET
					read = 'true',
					updated_at = NOW()
			WHERE user_id = $1
		`
	)

	result, err := r.db.Exec(ctx, query, req.UserId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
