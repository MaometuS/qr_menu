package item_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (i *itemRepository) Search(ctx context.Context, search string, establishmentID int64) ([]models.Item, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return nil, errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, `select * from items where 
		(name like concat('%', $1::varchar, '%') or additional_info like concat('%', $1::varchar, '%')) and
		category_id in (select id from categories where menu_id in (select id from menus where establishment_id = $2))`, search, establishmentID)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Item])
}
