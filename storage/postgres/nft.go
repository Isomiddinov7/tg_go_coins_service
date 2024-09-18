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

type nftRepo struct {
	db *pgxpool.Pool
}

func NewNFTRepo(db *pgxpool.Pool) storage.NFTRepoI {
	return &nftRepo{
		db: db,
	}
}

func (r *nftRepo) Create(ctx context.Context, req *coins_service.CreateNFT) (resp *coins_service.NFT, err error) {
	var (
		nft_id = uuid.NewString()
		query  = ` INSERT INTO "nft"(
			"id",
			"nft_img",
			"comment",
			"user_id",
			"telegram_id"
		) VALUES($1, $2, $3, $4, $5)`
	)

	_, err = r.db.Exec(ctx, query,
		nft_id,
		req.NftImg,
		req.Comment,
		req.UserId,
		req.TelegramId,
	)
	if err != nil {
		return nil, err
	}
	return r.GetById(ctx, &coins_service.NFTPrimaryKey{Id: nft_id})
}

func (r *nftRepo) GetById(ctx context.Context, req *coins_service.NFTPrimaryKey) (*coins_service.NFT, error) {
	var (
		query = `
			SELECT 
				"id",
				"nft_img",
				"comment",
				"user_id",
				"status",
				"telegram_id",
				"created_at",
				"updated_at"
			FROM "nft"
			WHERE "id" = $1
		`

		id          sql.NullString
		nft_img     sql.NullString
		comment     sql.NullString
		user_id     sql.NullString
		status      sql.NullString
		telegram_id sql.NullString
		created_at  sql.NullString
		updated_at  sql.NullString
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&nft_img,
		&comment,
		&user_id,
		&status,
		&telegram_id,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &coins_service.NFT{
		Id:         id.String,
		NftImg:     nft_img.String,
		Comment:    comment.String,
		UserId:     user_id.String,
		Status:     status.String,
		TelegramId: telegram_id.String,
		CreatedAt:  created_at.String,
		UpdatedAt:  updated_at.String,
	}, nil

}

func (r *nftRepo) GetAll(ctx context.Context, req *coins_service.GetListNFTRequest) (*coins_service.GetListNFTResponse, error) {
	var (
		resp   coins_service.GetListNFTResponse
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

	query := `
		SELECT
			COUNT(*) OVER(),
			"id",
			"nft_img",
			"comment",
			"user_id",
			"status",
			"telegram_id",
			"created_at",
			"updated_at"
		FROM "nft"
		`

	query += where + sort + offset + limit

	rowsNFT, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rowsNFT.Close()

	for rowsNFT.Next() {
		var (
			nft         coins_service.NFT
			id          sql.NullString
			nft_img     sql.NullString
			comment     sql.NullString
			user_id     sql.NullString
			status      sql.NullString
			telegram_id sql.NullString
			created_at  sql.NullString
			updated_at  sql.NullString
		)

		err = rowsNFT.Scan(
			&resp.Count,
			&id,
			&nft_img,
			&comment,
			&user_id,
			&status,
			&telegram_id,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}

		nft = coins_service.NFT{
			Id:         id.String,
			NftImg:     nft_img.String,
			Comment:    comment.String,
			UserId:     user_id.String,
			Status:     status.String,
			TelegramId: telegram_id.String,
			CreatedAt:  created_at.String,
			UpdatedAt:  updated_at.String,
		}

		resp.Nfts = append(resp.Nfts, &nft)
	}
	return &resp, nil
}

func (r *nftRepo) Update(ctx context.Context, req *coins_service.UpdateNFT) (int64, error) {
	var (
		query = `
		UPDATE "nft"
			SET
				"status" = $2,
				"updated_at" = NOW()
		WHERE "id" = $1`
	)

	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.Status,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (r *nftRepo) Delete(ctx context.Context, req *coins_service.NFTPrimaryKey) error {
	_, err := r.db.Exec(ctx, `DELETE FROM "nft" WHERE "id" = $1`, req.Id)
	if err != nil {
		return err
	}
	return nil
}
