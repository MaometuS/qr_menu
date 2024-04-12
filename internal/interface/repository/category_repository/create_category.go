package category_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (c *categoryRepository) CreateCategory(ctx context.Context, category *models.Category) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(
		ctx,
		"update categories set order_param=order_param+1 where order_param>=$1 and menu_id = $2",
		category.OrderParam,
		category.MenuID,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		ctx,
		"insert into categories(name, background, menu_id, is_visible, order_param) values ($1, $2, $3, $4, $5)",
		category.Name,
		category.Background,
		category.MenuID,
		category.IsVisible,
		category.OrderParam,
	)
	if err != nil {
		return err
	}

	return nil
}
