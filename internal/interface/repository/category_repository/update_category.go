package category_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (c *categoryRepository) UpdateCategory(ctx context.Context, category *models.Category) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(
		ctx,
		"update categories set name = $1, menu_id = $2, is_visible = $3, order_param = $4 where id = $5",
		category.Name,
		category.MenuID,
		category.IsVisible,
		category.OrderParam,
		category.ID,
	)
	if err != nil {
		return err
	}

	if category.Background != "" {
		_, err = db.Exec(ctx, "update categories set background = $1 where id = $2", category.Background, category.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
