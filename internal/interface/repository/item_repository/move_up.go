package item_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (i *itemRepository) MoveUp(ctx context.Context, itemID int64) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	var order, categoryID int64
	err := db.QueryRow(ctx, "select order_param, category_id from items where id = $1", itemID).Scan(&order, &categoryID)
	if err != nil {
		return err
	}

	if order == 1 {
		return errors.New("already first")
	}

	_, err = db.Exec(ctx, "update items set order_param=order_param+1 where order_param = $1 and category_id = $2", order-1, categoryID)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, "update items set order_param=order_param-1 where id = $1", itemID)
	if err != nil {
		return err
	}

	return nil
}
