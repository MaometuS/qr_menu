package category_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (c *categoryRepository) CheckBelongs(ctx context.Context, categoryID, profileID int64) (bool, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return false, errors.New("db not found in context")
	}

	var cnt int64
	err := db.QueryRow(
		ctx,
		`select count(id) from categories where id = $1 and menu_id in (
				select id from menus where establishment_id in (
					select id from establishments where profile_id = $2
				)
			 )`,
		categoryID,
		profileID,
	).Scan(&cnt)
	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}
