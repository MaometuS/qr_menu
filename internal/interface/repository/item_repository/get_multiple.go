package item_repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (i *itemRepository) GetMultiple(ctx context.Context, ids []int64) ([]models.Item, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return nil, errors.New("db not found in context")
	}

	if len(ids) == 0 {
		return []models.Item{}, nil
	}

	var t string
	for i := range ids {
		t += fmt.Sprintf("%d", ids[i])
		if i+1 < len(ids) {
			t += ", "
		}
	}
	t = "(" + t + ")"

	rows, err := db.Query(ctx, "select * from items where id in "+t)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Item])
}
