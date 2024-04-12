package item_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (i *itemRepository) GetEstablishment(ctx context.Context, itemID int64) (int64, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return 0, errors.New("db not found in context")
	}

	var establishmentID int64
	err := db.QueryRow(
		ctx,
		`select menus.establishment_id from items
				left join categories on categories.id = items.category_id
				left join menus on menus.id = categories.menu_id
            	group by menus.establishment_id, items.id
				having items.id = $1`,
		itemID,
	).Scan(&establishmentID)
	if err != nil {
		return 0, err
	}

	return establishmentID, nil
}
