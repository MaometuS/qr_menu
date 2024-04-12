package category_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (c *categoryRepository) GetOne(ctx context.Context, categoryID int64) (*models.Category, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return nil, errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, "select * from categories where id = $1", categoryID)
	if err != nil {
		return nil, err
	}

	category, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		return nil, err
	}

	return &category, nil
}
