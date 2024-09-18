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

type coinNftRepo struct {
	db *pgxpool.Pool
}

func NewCoinNftRepo(db *pgxpool.Pool) storage.CoinNFTRepoI {
	return &coinNftRepo{
		db: db,
	}
}

func (r *coinNftRepo) Create(ctx context.Context, req *coins_service.CoinNFTCreate) (resp *coins_service.CoinNFT, err error) {
	var (
		id    = uuid.NewString()
		query = ` INSERT INTO "coin_nft" (
			"id",
			"nft_img",
			"nft_price",
			"nft_address",
			"nft_name"
		) VALUES($1, $2, $3, $4, $5)`
	)

	_, err = r.db.Exec(ctx,
		query,
		id,
		req.NftImg,
		req.NftPrice,
		req.NftAddress,
		req.NftName,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, &coins_service.CoinNFTPrimaryKey{Id: id})
}

func (r *coinNftRepo) GetById(ctx context.Context, req *coins_service.CoinNFTPrimaryKey) (*coins_service.CoinNFT, error) {
	queryCoinNFT := `
		SELECT 
			"id",
			"nft_img",
			"nft_price",
			"nft_address",
			"nft_name",
			"created_at",
			"updated_at"
		FROM "coin_nft"
		WHERE id = $1
	`

	var (
		id          sql.NullString
		nft_img     sql.NullString
		nft_price   sql.NullString
		nft_address sql.NullString
		nft_name    sql.NullString
		created_at  sql.NullString
		updated_at  sql.NullString
	)

	err := r.db.QueryRow(ctx, queryCoinNFT, req.Id).Scan(
		&id,
		&nft_img,
		&nft_price,
		&nft_address,
		&nft_name,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &coins_service.CoinNFT{
		Id:         id.String,
		NftImg:     nft_img.String,
		NftPrice:   nft_price.String,
		NftAddress: nft_address.String,
		NftName:    nft_name.String,
		CreatedAt:  created_at.String,
		UpdatedAt:  updated_at.String,
	}, nil
}

func (r *coinNftRepo) GetList(ctx context.Context, req *coins_service.GetListCoinNFTRequest) (*coins_service.GetListCoinNFTResponse, error) {
	var (
		resp   coins_service.GetListCoinNFTResponse
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
		where += " AND nft_name ILIKE" + " '%" + req.Search + "%'"
	}

	query := `
		SELECT 
			COUNT(*) OVER(),
			"id",
			"nft_img",
			"nft_price",
			"nft_address",
			"nft_name",
			"created_at",
			"updated_at"
		FROM "coin_nft"
	`
	query += where + sort + offset + limit

	rowsCoinNFT, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rowsCoinNFT.Close()

	for rowsCoinNFT.Next() {
		var (
			coinNft     coins_service.CoinNFT
			id          sql.NullString
			nft_img     sql.NullString
			nft_price   sql.NullString
			nft_address sql.NullString
			nft_name    sql.NullString
			created_at  sql.NullString
			updated_at  sql.NullString
		)

		err = rowsCoinNFT.Scan(
			&resp.Count,
			&id,
			&nft_img,
			&nft_price,
			&nft_address,
			&nft_name,
			&created_at,
			&updated_at,
		)

		if err != nil {
			return nil, err
		}

		coinNft = coins_service.CoinNFT{
			Id:         id.String,
			NftImg:     nft_img.String,
			NftPrice:   nft_price.String,
			NftAddress: nft_address.String,
			NftName:    nft_name.String,
			CreatedAt:  created_at.String,
			UpdatedAt:  updated_at.String,
		}
		resp.CoinNfts = append(resp.CoinNfts, &coinNft)
	}
	return &resp, nil

}

func (r *coinNftRepo) Update(ctx context.Context, req *coins_service.CoinNFTUpdate) (int64, error) {
	var (
		query = `
			UPDATE "coin_nft"
			SET
				"nft_img" = $2,
				"nft_price" = $3,
				"nft_address" = $4,
				"nft_name" = $5,
				"updated_at" = NOW()
			WHERE "id" = $1`
	)

	rowsAffected, err := r.db.Exec(ctx,
		query,
		req.Id,
		req.NftImg,
		req.NftPrice,
		req.NftAddress,
		req.NftName,
	)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (r *coinNftRepo) Delete(ctx context.Context, req *coins_service.CoinNFTPrimaryKey) error {
	_, err := r.db.Exec(ctx, `DELETE FROM "coin_nft" WHERE "id" = $1`, req.Id)
	return err
}
