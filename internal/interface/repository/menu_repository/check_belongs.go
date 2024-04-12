package menu_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (m *menuRepository) CheckBelongs(ctx context.Context, menuID, profileID int64) (bool, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return false, errors.New("db not found in context")
	}

	var cnt int64
	err := db.QueryRow(
		ctx,
		"select count(id) from menus where id = $1 and establishment_id in (select id from establishments where profile_id = $2)",
		menuID,
		profileID,
	).Scan(&cnt)
	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}
