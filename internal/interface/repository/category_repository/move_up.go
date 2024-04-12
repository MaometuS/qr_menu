package category_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (c *categoryRepository) MoveUp(ctx context.Context, categoryID int64) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	var order, menuID int64
	err := db.QueryRow(ctx, "select order_param, menu_id from categories where id = $1", categoryID).Scan(&order, &menuID)
	if err != nil {
		return err
	}

	if order == 1 {
		return errors.New("already first")
	}

	_, err = db.Exec(ctx, "update categories set order_param=order_param+1 where order_param = $1 and menu_id = $2", order-1, menuID)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, "update categories set order_param=order_param-1 where id = $1", categoryID)
	if err != nil {
		return err
	}

	return nil
}
