package menu_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (m *menuRepository) GetAll(ctx context.Context, establishmentID int64) ([]models.Menu, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return nil, errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, "select * from menus where establishment_id = $1 order by order_param", establishmentID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Menu])
}
