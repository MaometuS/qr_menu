package establishment_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (e *establishmentRepository) GetByLink(ctx context.Context, link string) (*models.Establishment, error) {
	db, ok := ctx.Value("db").(entity.PgxIface)
	if !ok {
		return nil, errors.New("db not found in context")
	}

	rows, err := db.Query(ctx, "select * from establishments where link = $1", link)
	if err != nil {
		return nil, err
	}

	establishment, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Establishment])
	if err != nil {
		return nil, err
	}

	return &establishment, nil
}
