package item_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (i *itemRepository) UpdateItem(ctx context.Context, item *models.Item) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(
		ctx,
		"update items set name = $1, weight = $2, additional_info = $3, is_visible = $4, is_available = $5, price = $6, category_id = $7, order_param = $8 where id = $9",
		item.Name,
		item.Weight,
		item.AdditionalInfo,
		item.IsVisible,
		item.IsAvailable,
		item.Price,
		item.CategoryID,
		item.OrderParam,
		item.ID,
	)
	if err != nil {
		return err
	}

	if item.Image != "" {
		_, err = db.Exec(ctx, "update items set image = $1 where id = $2", item.Image, item.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
