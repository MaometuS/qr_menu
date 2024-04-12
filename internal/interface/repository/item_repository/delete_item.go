package item_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (i *itemRepository) DeleteItem(ctx context.Context, itemID int64) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, "select * from items where id = $1", itemID)
	if err != nil {
		return err
	}

	item, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Item])
	if err != nil {
		return err
	}

	_, err = db.Exec(
		ctx,
		"update items set order_param=order_param-1 where order_param>$1 and category_id = $2",
		item.OrderParam,
		item.CategoryID,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, "delete from items where id = $1", itemID)
	if err != nil {
		return err
	}

	return nil
}
