package item_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (i *itemRepository) CheckBelongs(ctx context.Context, itemID, profileID int64) (bool, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return false, errors.New("db not found in context")
	}

	var cnt int64
	err := db.QueryRow(
		ctx,
		`select count(id) from items where id = $1 and category_id in (
    			select id from categories where menu_id in (
    			    select id from menus where establishment_id in (
    			        select id from establishments where profile_id = $2
					)
				)
			 )`,
		itemID,
		profileID,
	).Scan(&cnt)
	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}
