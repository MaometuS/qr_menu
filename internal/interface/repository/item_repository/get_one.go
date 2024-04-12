package item_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (i *itemRepository) GetOne(ctx context.Context, itemID int64) (*models.Item, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return nil, errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, "select * from items where id = $1", itemID)
	if err != nil {
		return nil, err
	}

	item, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Item])
	if err != nil {
		return nil, err
	}

	return &item, nil
}
