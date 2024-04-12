package category_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (c *categoryRepository) DeleteCategory(ctx context.Context, categoryID int64) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, "select * from categories where id = $1", categoryID)
	if err != nil {
		return err
	}

	category, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		return err
	}

	_, err = db.Exec(
		ctx,
		"update categories set order_param=order_param-1 where order_param>$1 and menu_id = $2",
		category.OrderParam,
		category.MenuID,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, "delete from categories where id = $1", categoryID)
	if err != nil {
		return err
	}

	return nil
}
