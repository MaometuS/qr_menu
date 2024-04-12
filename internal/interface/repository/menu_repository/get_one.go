package menu_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (m *menuRepository) GetOne(ctx context.Context, menuID int64) (*models.Menu, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return nil, errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, "select * from menus where id = $1", menuID)
	if err != nil {
		return nil, err
	}

	menu, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Menu])
	if err != nil {
		return nil, err
	}

	return &menu, nil
}
