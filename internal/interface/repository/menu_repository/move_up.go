package menu_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (m *menuRepository) MoveUp(ctx context.Context, menuID int64) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	var order, establishmentID int64
	err := db.QueryRow(ctx, "select order_param, establishment_id from menus where id = $1", menuID).Scan(&order, &establishmentID)
	if err != nil {
		return err
	}

	if order == 1 {
		return errors.New("already first")
	}

	_, err = db.Exec(ctx, "update menus set order_param=order_param+1 where order_param = $1 and establishment_id = $2", order-1, establishmentID)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, "update menus set order_param=order_param-1 where id = $1", menuID)
	if err != nil {
		return err
	}

	return nil
}
