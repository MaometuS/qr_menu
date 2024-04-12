package menu_repository

import (
	"context"
	"errors"
	"gitlab.com/maometusu/qr_menu/internal/entity"
)

func (m *menuRepository) GetTopID(ctx context.Context, establishmentID int64) (int64, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return 0, errors.New("db not found in context")
	}

	var topID int64
	err := db.QueryRow(ctx, "select id from menus where establishment_id = $1 order by order_param limit 1", establishmentID).Scan(&topID)
	if err != nil {
		return 0, err
	}

	return topID, nil
}
