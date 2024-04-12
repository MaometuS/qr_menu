package profile_repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/entity/models"
)

func (p *profileRepository) GetOne(context context.Context, id int64) (*models.Profile, error) {
	db, ok := context.Value("db").(entity.PgxIface)
	if !ok {
		return nil, errors.New("connection is missing from context")
	}

	rows, err := db.Query(context, "select * from profiles where id = $1", id)
	if err != nil {
		return nil, err
	}

	profile, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Profile])
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
