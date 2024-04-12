package menu_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (m *menuRepository) DeleteMenu(ctx context.Context, menuID int64) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, "select * from menus where id = $1", menuID)
	if err != nil {
		return err
	}

	menu, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Menu])
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, "update menus set order_param=order_param-1 where order_param > $1 and establishment_id = $2", menu.OrderParam, menu.EstablishmentID)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, "delete from menus where id = $1", menuID)
	if err != nil {
		return err
	}

	return nil
}
