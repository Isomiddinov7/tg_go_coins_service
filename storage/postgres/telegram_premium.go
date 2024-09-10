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

type premiumRepo struct {
	db *pgxpool.Pool
}

func NewTelegramPremiumRepo(db *pgxpool.Pool) storage.TelegramPremiumRepoI {
	return &premiumRepo{
		db: db,
	}
}

func (r *premiumRepo) CreatePrice(ctx context.Context, req *coins_service.CreateTelegramPremiumPrice) (resp *coins_service.TelegramPremiumPrice, err error) {
	var (
		query = `
			INSERT INTO "premium_price_month"(
				"id",
				"month",
				"price",
				"premium_id"
			) VALUES($1, $2, $3, $4)
		`
		id = uuid.NewString()
	)

	_, err = r.db.Exec(ctx,
		query,
		id,
		req.Month,
		req.Price,
		req.PremiumId,
	)
	if err != nil {
		return nil, err
	}

	return &coins_service.TelegramPremiumPrice{
		Id:    id,
		Month: req.Month,
		Price: req.Price,
	}, nil
}

func (r *premiumRepo) CreatePremium(ctx context.Context, req *coins_service.CreateTelegramPremium) (resp *coins_service.TelegramPremium, err error) {
	var (
		query = `
			INSERT INTO "premium"(
				"id",
				"name",
				"card_number",
				"img"
			) VALUES($1, $2, $3, $4)
		`
		id = uuid.NewString()
	)

	_, err = r.db.Exec(ctx, query,
		id,
		req.Name,
		req.CardNumber,
		req.Img,
	)
	if err != nil {
		return nil, err
	}
	return r.GetPremiumById(ctx, &coins_service.TelegramPriemiumPrimaryKey{Id: id})
}

func (r *premiumRepo) GetPremiumById(ctx context.Context, req *coins_service.TelegramPriemiumPrimaryKey) (resp *coins_service.TelegramPremium, err error) {
	var (
		query = `
			SELECT
				"id",
				"name",
				"card_name",
				"img",
				"created_at",
				"updated_at"
			FROM "premium"
			WHERE id = $1
		`

		id         sql.NullString
		name       sql.NullString
		card_name  sql.NullString
		img        sql.NullString
		prices     []*coins_service.TelegramPremiumPrice
		created_at sql.NullString
		updated_at sql.NullString

		queryPrice = `
			SELECT 
				"id",
				"month",
				"price"
			FROM "premium_price_month"
			WHERE premium_id = $1
		`
	)

	err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&card_name,
		&img,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, queryPrice, req.Id)

	for rows.Next() {
		var (
			data     coins_service.TelegramPremiumPrice
			price_id sql.NullString
			month    sql.NullString
			price    sql.NullString
		)

		err = rows.Scan(
			&price_id,
			&month,
			&price,
		)
		data = coins_service.TelegramPremiumPrice{
			Id:    price.String,
			Month: month.String,
			Price: price.String,
		}

		prices = append(prices, &data)
	}

	return &coins_service.TelegramPremium{
		Id:         id.String,
		Name:       name.String,
		CardNumber: card_name.String,
		Img:        img.String,
		Price:      prices,
		CreatedAt:  created_at.String,
		UpdatedAt:  updated_at.String,
	}, nil
}

func (r *premiumRepo) UpdateTransactionStatus(ctx context.Context, req *coins_service.UpdateStatus) (int64, error) {
	var (
		query = `
			UPDATE "premium_transaction"
				SET
					"transaction_status" = $2
					"updated_at" = NOW()
				WHERE "id" = $1
		`
	)

	rowsAffected, err := r.db.Exec(ctx, query, req.TransactionId, req.TransactionStatus)
	if err != nil {
		return 0, err
	}
	return rowsAffected.RowsAffected(), nil
}

func (r *premiumRepo) PremiumTransaction(ctx context.Context, req *coins_service.PremiumTransactionRequest) error {
	var (
		query = `
			INSERT INTO "premium_transaction"(
				"id",
				"phone_number",
				"telegram_username",
				"price_id",
				"payment_img",
				"user_id"
			) VALUES($1, $2, $3, $4, $5, $6)
		`

		id = uuid.NewString()
	)

	_, err := r.db.Exec(ctx,
		query,
		id,
		req.PhoneNumber,
		req.TelegramUsername,
		req.TelegramPriceId,
		req.PaymentImg,
		req.UserId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *premiumRepo) GetList(ctx context.Context, req *coins_service.GetListPremiumRequest) (*coins_service.GetPremiumTransactionResponse, error) {
	var (
		resp   coins_service.GetPremiumTransactionResponse
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

	var (
		query = `
			SELECT
				COUNT(*) OVER(),
				pt."telegram_username",
				pm."name",
				pm."month",
				u."first_name",
				pt."created_at",
				pt."updated_at"
			FROM "premium_transaction" as pt
			JOIN "premium" as p ON p."id" = pt."premium_id"
			JOIN "users" as u ON u."id" = pt."user_id"
			JOIN "premium_price_month" as pm ON pm."premium_id"= p."id"
		`
	)
	query += where + sort + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			data              coins_service.GetPremiumTransaction
			telegram_username sql.NullString
			name              sql.NullString
			month             sql.NullString
			first_name        sql.NullString
			created_at        sql.NullString
			updated_at        sql.NullString
		)

		err = rows.Scan(
			resp.Count,
			&telegram_username,
			&name,
			&month,
			&first_name,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}

		data = coins_service.GetPremiumTransaction{
			UserName:  telegram_username.String,
			Name:      name.String,
			Month:     month.String,
			FirstName: first_name.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		}

		resp.Transactions = append(resp.Transactions, &data)
	}
	return &resp, nil
}
