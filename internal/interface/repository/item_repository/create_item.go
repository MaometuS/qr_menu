package item_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (i *itemRepository) CreateItem(ctx context.Context, item *models.Item) error {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return errors.New("db not found in context")
	}

	_, err := db.Exec(
		ctx,
		"update items set order_param=order_param+1 where order_param>=$1 and category_id = $2",
		item.OrderParam,
		item.CategoryID,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		ctx,
		"insert into items(name, weight, additional_info, is_visible, is_available, image, category_id, order_param, price) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		item.Name,
		item.Weight,
		item.AdditionalInfo,
		item.IsVisible,
		item.IsAvailable,
		item.Image,
		item.CategoryID,
		item.OrderParam,
		item.Price,
	)
	if err != nil {
		return err
	}

	return nil
}
