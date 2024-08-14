package postgres

import (
	"context"
	"tg_go_coins_service/genproto/coins_service"
	"tg_go_coins_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type imageRepo struct {
	db *pgxpool.Pool
}

func NewFileImageRepo(db *pgxpool.Pool) storage.ImagesRepoI {
	return &imageRepo{
		db: db,
	}
}

func (r *imageRepo) FileUpload(ctx context.Context, req *coins_service.ImageData) (resp *coins_service.ImagePrimaryKey, err error) {

	var (
		query = `
			INSERT INTO "images"("id", "image_link") VALUES($1, $2)
		`
	)

	_, err = r.db.Exec(ctx, query, req.Id, req.ImageLink)
	if err != nil {
		return nil, err
	}

	return &coins_service.ImagePrimaryKey{
		Id: req.Id,
	}, nil
}

func (r *imageRepo) FileDelete(ctx context.Context, req *coins_service.ImagePrimaryKey) (resp *coins_service.Empty, err error) {
	_, err = r.db.Exec(ctx, `DELETE FROM "images" WHERE "id" = $1`, req.Id)
	return nil, err
}
